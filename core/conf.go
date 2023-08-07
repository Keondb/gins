package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "settings.yaml"

func InitConf() {
	//const ConfigFile = "settings.yaml"

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

func Setyaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
