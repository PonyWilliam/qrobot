package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf struct {
	AppID  uint64 `yaml:"appid"`
	Token  string `yaml:"token"`
	Secret string `yaml:"secret"`
}

func catcherr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func init() {
	//将配置读取到内存中

	content, err := ioutil.ReadFile("./config.yaml")
	catcherr(err)
	err = yaml.Unmarshal(content, &Conf)
	catcherr(err)
}
