package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func LoadConfigMap(fileName string) interface{} {
	var config interface{}

	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Panic("Open config file fail", err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panic("call ioutil fail", err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic("parse json fail", err)
	}

	return config
}

func FillStruct(data map[string]interface{}, result *interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}

func LoadConfig(fileName string, result *interface{}) {
	data := LoadConfigMap(fileName).(map[string]interface{})
	FillStruct(data, result)
}
