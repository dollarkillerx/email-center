/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:23 2019-09-25
 */
package config

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type myconf struct {
	App struct {
		Host       string `yaml:"host"`
		Debug      bool   `yaml:"debug"`
		MaxRequest int    `yaml:"max_request"`
	} `yaml:"app"`
	Email []Email `yaml:"email"`
}

type Email struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

var (
	MyConfig *myconf
)

func init() {
	MyConfig = &myconf{}

	b, i := easyutils.PathExists("./config.yml")
	if i != nil || b == false {
		create()
	}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}

	if MyConfig.App.Debug {
		log.Println(MyConfig)
	}
}
func create() {
	err := ioutil.WriteFile("./config.yml", []byte(emconf), 00666)
	if err != nil {
		panic("文件创建失败")
	}
	clog.Println("请填写配置文件")
	os.Exit(0)
}

var emconf = `
# 通用配置文件

app:
  host: "0.0.0.0:8083"
  debug: true
  max_request: 1000 # 最大并发数量

email: # email 集合
  - user:
    password:
    host: "smtp.mail.ru"
    port: 465
`
