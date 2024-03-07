package main

import (
	"context"
	"encoding/json"
	"fmt"
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
}
func (a *App) shutdown(ctx context.Context) {
	global.GlobalConfig.SaveConf()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetConfig() string {
	jsonData, err := json.Marshal(global.GlobalConfig)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
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
	res, err := http.Get(fmt.Sprintf(global.SEARCHURL, pageNo, keyWord))
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	var result model.SearchRes
	// js要解析有些地方还得加转义字符，先在后端转换一下
	json.NewDecoder(res.Body).Decode(&result)
	bytes, _ := json.Marshal(result)
	return string(bytes)
}
