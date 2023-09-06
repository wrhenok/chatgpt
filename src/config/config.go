package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
)

var Server *server
var Database *database
var ApiKey string
var Url *url

type conf struct {
	Svc    server   `yaml:"server"`
	DB     database `yaml:"database"`
	Apikey string   `yaml:"apikey"`
	Url    url      `yaml:"url"`
}

type database struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

type server struct {
	Port string `yaml:"port"`
}

type url struct {
	BaseUrl   string `ymal:"baseurl"`
	SuffixUrl string `ymal:"suffixurl"`
	UrlSuffix string `yaml:"urlsuffix"`
}

func InitConf(dataFile string) {
	// 解决相对路经下获取不了配置文件问题
	_, filename, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(filename), dataFile)
	_, err := os.Stat(filePath)
	if err != nil {
		log.Printf("config file path %s not exist", filePath)
	}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(conf)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
	// 绑定到外部可以访问的变量中
	Server = &c.Svc
	Database = &c.DB
	ApiKey = c.Apikey
	Url = &c.Url
	println(Url.BaseUrl)
}
