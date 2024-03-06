package config

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/spf13/viper"
)

type NaCosConf struct {
	Ip     string `yaml:"Ip"`
	Port   int    `yaml:"Port"`
	DataId string `yaml:"DataId"`
	Group  string `yaml:"Group"`
}
type Dial struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}
type JwtConf struct {
	Password string `yaml:"Password"`
}

var N NaCosConf
var D Dial
var J JwtConf

func init() {
	viper.SetConfigFile("/Users/jiangshiao/go/src/zg5/work/framework/config/test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logs.Info(err, "viper")
		return
	}
	N.Ip = viper.GetString("NaCos.Ip")
	N.Port = viper.GetInt("NaCos.Port")
	N.DataId = viper.GetString("NaCos.DataId")
	N.Group = viper.GetString("NaCos.Group")
	fmt.Println(N, "nacos 配置信息")
	D.Host = viper.GetString("Dial.Host")
	D.Port = viper.GetString("Dial.Port")
	fmt.Println(D)
	J.Password = viper.GetString("Jwt.Password")
}
