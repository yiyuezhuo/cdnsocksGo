package main

import (
	"fmt"

	"github.com/yiyuezhuo/cdnsocksGo/utils"
)

type Route struct {
	Name      string
	LocalIp   string
	LocalPort string
	Url       string
}

type Config struct {
	RemoteIp   string
	RemotePort string
	Routes     []Route
}

func (config Config) Print() {
	fmt.Println("RemoteIp:", config.RemoteIp)
	fmt.Println("RemotePort", config.RemotePort)
	for idx, value := range config.Routes {
		fmt.Println(" Idx:", idx)
		fmt.Println(" Name:", value.Name)
		fmt.Println(" LocalIp:", value.LocalIp)
		fmt.Println(" LocalPort:", value.LocalPort)
		fmt.Println(" Url:", value.Url)
		fmt.Println("")
	}
}

func LoadConfig(fileName string) Config {
	var config Config
	utils.LoadConfig(fileName, &config)
	return config
}
