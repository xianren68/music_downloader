package main

import (
	"context"
	"encoding/json"
	"os"

	"music_downloader/global"
	"music_downloader/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v3"
)

func init() {
	// 判断当前文件夹是否有配置文件
	_, err := os.Stat("./config.yaml")
	if os.IsNotExist(err) {
		// 如果没有配置文件，使用默认配置
		global.GlobalConfig = global.DefaultConfig()
		// 创建并打开文件
		file, err := os.Create("./config.yaml")
		if err != nil {
			os.Exit(1)
		}
		defer file.Close()
		// 将配置文件写入文件
		data, err := yaml.Marshal(global.GlobalConfig)
		if err != nil {
			os.Exit(1)
		}
		os.WriteFile("./config.yaml", data, 0644)
		return
	}
	// 配置文件存在，则加载配置文件
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		os.Exit(1)
	}
	// 解析配置文件
	err = yaml.Unmarshal(data, global.GlobalConfig)
	if err != nil {
		os.Exit(1)
	}
}

// readTemp 获取缓存文件.
func readTemp() {
	var ctx = context.Background()
	// 缓存文件是否存在
	_, err := os.Stat("./temp.json")
	if os.IsNotExist(err) {
		return
	}
	// 读取缓存文件
	data, err := os.ReadFile("./temp.json")
	if err != nil {
		runtime.EventsEmit(ctx, "warning", "读取缓存文件失败")
		return
	}
	var tempData []*model.Music
	// 解析缓存文件
	err = json.Unmarshal(data, &tempData)
	if err != nil {
		runtime.EventsEmit(ctx, "warning", "解析缓存文件失败")
		return
	}
	runtime.EventsEmit(ctx, "info", "继续下载未完成任务")
	// 更新全局变量
	for _, v := range tempData {
		model.Down.NewTask(v)
	}
	runtime.EventsEmit(ctx, "continueDownload", model.Down.ExportList())
}
