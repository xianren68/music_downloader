package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Package Config 总配置文件.
type Config struct {
	Save *SaveConfig `yaml:"save"`
}

// 保存配置
func (c *Config) SaveConf() {
	data, _ := yaml.Marshal(c)
	_ = os.WriteFile("config.yaml", data, 0644)
}
