package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func InitConf() {
	const ConfigFile = "settings.yaml"

	conf := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error : %s", err))
	}

	err = yaml.Unmarshal(yamlConf, conf)
	if err != nil {
		log.Fatalf("config Init Unmarshal : %v", err)
	}
	log.Println("config sucess")
	// fmt.Println(conf)

	global.Config = conf
}
