package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config is
type Config struct {
	Gdas Gdas `yaml:"gdas"`
	Xsky Xsky `yaml:"xsky"`
}

var yamlConfig Config

// Gdas is
type Gdas struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Path     string `yaml:"path"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Xsky is
type Xsky struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Path     string `yaml:"path"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// GetConfig 用来获取 yaml 格式配置信息
func init() {
	file, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("获取配置文件错误：%v", err)
	}

	if err = yaml.Unmarshal(file, &yamlConfig); err != nil {
		log.Fatalf("解析配置文件错误：%v", err)
	}
	// fmt.Println(yamlConfig)
}
