package main

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	var config Config
	LoadConfig("config-client.json", config)
	config.Print()
}
