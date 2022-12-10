package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MysqlURI    string      `yaml:"mysqlURI"`
	CasdoorConf CasdoorConf `yaml:"casdoorConf"`
	MinioConf   MinioConf   `yaml:"minioConf"`
}

type CasdoorConf struct {
	Endpoint            string `yaml:"endpoint"`
	ClientId            string `yaml:"clientId"`
	ClientSecret        string `yaml:"clientSecret"`
	JwtSecret           string `yaml:"jwtSecret"`
	CasdoorOrganization string `yaml:"casdoorOrganization"`
	CasdoorApplication  string `yaml:"casdoorApplication"`
}

type MinioConf struct {
	Endpoint     string `yaml:"endpoint"`
	AccessKey    string `yaml:"accessKey"`
	AccessSecret string `yaml:"accessSecret"`
	UseSSL       bool   `yaml:"useSSL"`
}

var Conf *Config

func InitConfig() {
	Conf = &Config{}
	var (
		err  error
		conf []byte
	)
	if os.Getenv("ENV") == "prod" {
		conf, err = os.ReadFile("./conf/echo-bio-cloud.yaml")
	} else {
		conf, err = os.ReadFile("./conf/echo-bio-cloud-dev.yaml")
	}
	if err != nil {
		log.Fatal(fmt.Sprint("read config file failed:", err))
	}
	err = yaml.Unmarshal(conf, Conf)
	if err != nil {
		log.Fatal("parse config file failed:", err)
	}
}
