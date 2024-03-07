package main

import (
	"os"

	"music_downloader/global"

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
