// Package global 一些全局变量.
package global

import (
	"music_downloader/config"
	"path/filepath"
)

var GlobalConfig *config.Config = new(config.Config)

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
