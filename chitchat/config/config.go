package config

import (
	"encoding/json"
	"os"
	"sync"
	"log"
)

type App struct {
	Address      string
	Static       string
	Log          string
}

type Database struct {
	Driver      string
	Address        string
	Database    string
	User        string
	Password    string
}

type Configuration struct {
	App App
	Db  Database
}

var configs *Configuration
var once sync.Once

// 通过单例模式初始化全局配置
func LoadConfig() *Configuration {
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		configs = &Configuration{}
		err = decoder.Decode(configs)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})
	return configs
}