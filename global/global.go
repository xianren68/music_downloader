// Package global 一些全局变量.
package global

import (
	"music_downloader/config"
	"path/filepath"
)

var GlobalConfig *config.Config = new(config.Config)

const SEARCHURL = "http://search.kuwo.cn/r.s?client=kt&all=%s&pn=%d&rn=20&uid=221260053&ver=kwplayer_ar_99.99.99.99&vipver=1&ft=music&cluster=0&strategy=2012&encoding=utf8&rformat=json&vermerge=1&mobi=1"
const DOWNLOADURL = "http://antiserver.kuwo.cn/anti.s?type=convert_url&rid=%s&format=mp3|acc&response=url"

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
