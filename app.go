package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"music_downloader/config"
	"music_downloader/global"
	"music_downloader/model"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.readTemp()
}

// 关闭软件生命周期钩子.
func (a *App) shutdown(ctx context.Context) {
	global.GlobalConfig.SaveConf()
}

func (a *App) onBeforeClose(ctx context.Context) bool {
	runtime.EventsEmit(a.ctx, "ending", nil)
	return false
}

// GetConfig 获取配置文件.
func (a *App) GetConfig() string {
	jsonData, err := json.Marshal(global.GlobalConfig)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// OpenDirDialog 选择要存储的文件.
func (a *App) OpenDirDialog() string {
	s, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: "D:\\",
	})
	return s
}
func (a *App) SaveConfig(confi string) bool {
	var conf config.Config
	err := json.Unmarshal([]byte(confi), &conf)
	// 保存配置失败
	if err != nil {
		return false
	}
	if conf.Save.Path == global.GlobalConfig.Save.Path && conf.Save.IsSortByAuthor == global.GlobalConfig.Save.IsSortByAuthor {
		return true
	}
	global.GlobalConfig = &conf
	return true

}

func (a *App) Search(pageNo int, keyWord string) string {
	res, err := http.Get(fmt.Sprintf(global.SEARCHURL, keyWord, pageNo))
	if err != nil {
		runtime.EventsEmit(a.ctx, "error", "请求出错")
		return ""
	}
	defer res.Body.Close()
	bytes, _ := io.ReadAll(res.Body)
	return string(bytes)
}

func (a *App) Download(name, author, id, album string) {
	music := &model.Music{
		Name:   name,
		Author: author,
		ID:     id,
		ALBUM:  album,
	}
	// 添加新任务
	model.Down.NewTask(a.ctx, music)

}
