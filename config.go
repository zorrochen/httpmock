package main

import (
	"github.com/jinzhu/configor"
	"log"
)

type Config_t struct {
	SrvAddr   	string `yaml:"SERVER_ADDR"`
	MockFile   	string `yaml:"MOCK_FILE"`
}

var CfgData *Config_t

func ConfigInit(confPath string) {
	var path string
	path = confPath

	var cfg = &Config_t{}
	if err := configor.Load(cfg, path); err != nil {
		panic(err)
	}
	CfgData = cfg

	log.Printf("[ConfigInit] %+v", CfgData)
}
