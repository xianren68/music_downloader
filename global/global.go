// Package global 一些全局变量.
package global

import (
	"music_downloader/config"
	"path/filepath"
)

var GlobalConfig *config.Config = new(config.Config)

const SEARCHURL = "http://pd.musicapp.migu.cn/MIGUM2.0/v1.0/content/search_all.do?ua=Android_migu&version=5.0.1&pageNo=%d&pageSize=10&text=%s&searchSwitch="

// DefaultConfig 默认配置.
func DefaultConfig() *config.Config {
	config := &config.Config{
		Save: &config.SaveConfig{
			IsSortByAuthor: false,
		},
	}
	config.Save.Path, _ = filepath.Abs("./download")
	return config
}
