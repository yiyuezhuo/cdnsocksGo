package main

import (
	"github.com/yiyuezhuo/cdnsocksGo/utils"
)

type Route struct {
	Name string
	Url  string
	Type string
	Port string
}

type Config struct {
	ListenIp   string
	ListenPort string
	Routes     []Route
}

func LoadConfig(filePath string) Config {
	return utils.LoadConfig(filePath).(Config)
}
