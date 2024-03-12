// Package model 定义数据模型.
package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"music_downloader/global"
	"net/http"
	"os"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Music 音乐.
type Music struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	ID     string `json:"id"`
	ALBUM  string `json:"album"`
}

// Node 链表节点.
type Node struct {
	Value *Music
	Next  *Node
	Prev  *Node
}

// Progress 返回下载进度.
type Progress struct {
	ID       string  `json:"id"`
	Status   bool    `json:"status"`
	Progress float64 `json:"progress"`
}

// List 双向链表.
type List struct {
	// 属性全部私有
	head *Node
	tail *Node
	size int
	// 映射，方便删除某个任务
	hash map[string]*Node
}

// AddNode 添加节点
func (l *List) AddNode(value *Music) {
	node := &Node{
		Value: value,
		Next:  l.tail,
		Prev:  nil,
	}
	node.Prev = l.tail.Prev
	l.tail.Prev.Next = node
	l.tail.Prev = node
	l.hash[value.ID] = node
	l.size++
}

// RemoveNode 删除指定节点
func (l *List) RemoveNode(ID string) {
	node, ok := l.hash[ID]
	if !ok {
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	delete(l.hash, ID)
	l.size--
}

// PopNode 弹出节点.
func (l *List) PopNode() *Node {
	if l.size == 0 {
		return nil
	}
	node := l.head.Next
	l.head.Next = node.Next
	node.Next.Prev = l.head
	l.size--
	delete(l.hash, node.Value.ID)
	return node
}

// Size 获取长度
func (l *List) Size() int {
	return l.size
}

// InitList 初始化链表.
func InitList() *List {
	l := &List{
		head: &Node{},
		tail: &Node{},
		hash: make(map[string]*Node),
		size: 0,
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

// Download 下载器
type Download struct {
	// 等待队列
	waitList *List
	// 是否开始下载任务
	isStart bool
	// 是否开启锁
	startLock sync.Mutex
	// 下载锁
	lock sync.Mutex
	// 缓存管道（用于保存再关闭时还未完成的任务）
	tempChan chan *Music
}

var Down = NewDownload()

// NewDownload 创建下载器
func NewDownload() *Download {
	return &Download{
		waitList:  InitList(),
		tempChan:  make(chan *Music, 5),
		startLock: sync.Mutex{},
		lock:      sync.Mutex{},
		isStart:   false,
	}
}

// NewTask 添加新的下载任务到队列
func (d *Download) NewTask(ctx context.Context, value *Music) {
	// 添加下载任务到队列
	d.lock.Lock()
	d.waitList.AddNode(value)
	d.lock.Unlock()
	d.startLock.Lock()
	// 加锁
	if !d.isStart {
		d.isStart = true
		// 开启控制下载协程
		go func() {
			// 任务队列，同时只能有4个协程下载
			task := make(chan struct{}, 4)
			wg := sync.WaitGroup{}
			for {
				// 加锁获取等待队列中的任务
				d.lock.Lock()
				if d.waitList.Size() == 0 {
					// 没有任务了
					d.lock.Unlock()
					// 等待剩余协程全部完成
					wg.Wait()
					// 是否有新任务出现
					d.lock.Lock()
					if d.waitList.Size() == 0 {
						d.lock.Unlock()
						// 没有则关闭控制下载协程
						break
					}
					d.lock.Unlock()
					continue
				}
				value := d.waitList.PopNode()
				d.lock.Unlock()
				// 任务添加到管道
				task <- struct{}{}
				// 等待到有空闲，执行下载任务
				go download(ctx, task, value.Value, d.tempChan, &wg)
			}
			// 修改状态
			d.startLock.Lock()
			d.isStart = false
			d.startLock.Unlock()
		}()
	}
	d.startLock.Unlock()
}

// ExportList 导出下载队列
func (d *Download) ExportList() []*Music {
	d.lock.Lock()
	defer d.lock.Unlock()
	var result []*Music
	for _, v := range d.waitList.hash {
		result = append(result, v.Value)
	}
	return result
}

func (d *Download) End() {
	// 关闭控制下载协程
	d.startLock.Lock()
	d.isStart = false
	d.startLock.Unlock()
	// 取出所有值
	tempJson := d.ExportList()
	d.lock.Lock()
	d.waitList.size = 0
	d.lock.Unlock()
label:
	for {
		select {
		case v, ok := <-d.tempChan:
			if ok {
				tempJson = append(tempJson, v)
			}
		default:
			break label
		}

	}
	if len(tempJson) > 0 {
		fmt.Println("下载完成，正在保存数据...")
		saveTemp(tempJson)
	}

}

// download 下载音乐.
func download(ctx context.Context, task chan struct{}, value *Music, temp chan *Music, wg *sync.WaitGroup) {
	// 执行前保存到缓存管道
	temp <- value
	// 添加一个等待协程
	wg.Add(1)
	defer wg.Done()
	url := fmt.Sprintf(global.DOWNLOADURL, value.ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		runtime.EventsEmit(ctx, "error", "请求出错")
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		runtime.EventsEmit(ctx, "error", "请求出错")
	}
	// 获取歌曲对应的url
	trueUrl, err := io.ReadAll(res.Body)
	if err != nil {
		runtime.EventsEmit(ctx, "error", "请求出错")
	}
	res, err = http.Get(string(trueUrl))
	if err != nil {
		runtime.EventsEmit(ctx, "error", "请求出错")
	}
	defer res.Body.Close()
	// 要保存的文件名
	fileName := value.Name + ".mp3"
	// 要保存的文件夹
	folder := global.GlobalConfig.Save.Path
	// 需要按作者分类
	if global.GlobalConfig.Save.IsSortByAuthor {
		folder = folder + "\\" + value.Author
	}
	_, err = os.Stat(folder)
	if os.IsNotExist(err) {
		// 创建文件夹
		os.Mkdir(folder, os.ModePerm)
	}
	orifile, err := os.Stat(folder + "\\" + fileName)
	// 请求长度
	leng := res.ContentLength
	if os.IsNotExist(err) || orifile.Size() != leng {
		// 创建文件
		file, _ := os.OpenFile(folder+"\\"+fileName, os.O_CREATE|os.O_RDWR, 0666)
		defer file.Close()
		// 写入文件
		bytes := [1024]byte{}
		i := 0
		var Pro = &Progress{
			ID: value.ID,
		}
		for {
			n, err := res.Body.Read(bytes[:])
			if n == 0 {
				Pro.Progress = 1
				Pro.Status = true
				runtime.EventsEmit(ctx, "downloadProgress", Pro)
				break
			}
			if err != nil && err != io.EOF {
				Pro.Progress = 1
				Pro.Status = true
				runtime.EventsEmit(ctx, "downloadProgress", Pro)
				break
			}
			if i != 0 && i%1024 == 0 {
				Pro.Progress = float64(i) / float64(leng/1024)
				Pro.Status = false
				runtime.EventsEmit(ctx, "downloadProgress", Pro)
			}
			file.Write(bytes[:n])
			i++
		}
	} else if orifile.Size() == leng {
		runtime.EventsEmit(ctx, "info", value.Name+"已存在")
	}

	// 任务完成，从缓存管道弹出值
	<-temp
	<-task
}

// saveTemp 存储缓存.
func saveTemp(value []*Music) {
	file, _ := os.Create("./temp.json")
	defer file.Close()
	data, _ := json.Marshal(value)
	file.Write(data)
}
