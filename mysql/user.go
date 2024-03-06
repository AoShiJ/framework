package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var T1 T

func InitMysql(string2 string) error {
	var err error
	json.Unmarshal([]byte(string2), &T1)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		T1.Mysql.Username,
		T1.Mysql.Password,
		T1.Mysql.Host,
		T1.Mysql.Port,
		T1.Mysql.Database,
	)
	logs.Info(dsn, "mysql 配置")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

type T struct {
	App struct {
		Ip   string `json:"Ip"`
		Port string `json:"Port"`
	} `json:"app"`
	Mysql struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     string `json:"Port"`
		Database string `json:"Database"`
	} `json:"Mysql"`
	Redis struct {
		Ip   string `json:"Ip"`
		Port int    `json:"Port"`
	} `json:"Redis"`
}
