package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 应用配置实例
var App = &AppConfig{}

type AppConfig struct {
	Server *ServerConfig `yaml:"server"`

	Logger *Logger `yaml:"logger"`

	Work *WorkConfig `yaml:"work"`

	Redis *RedisConfig `yaml:"redis"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

/**
 * @Description: 初始化配置
 */
func init() {
	data, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		panic("解析配置文件出错")
	}
	yaml.Unmarshal(data, App)
}
