// Package config 配置包.
package config

// Save 存储相关.
type SaveConfig struct {
	Path string `yaml:"Path"`
	// 下载文件是否按作者分类
	IsSortByAuthor bool `yaml:"isSortByAuthor"`
}
