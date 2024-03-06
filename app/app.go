package app

import (
	"github.com/AoShiJ/framework/mysql"
	"github.com/AoShiJ/framework/nacos"
	"github.com/AoShiJ/framework/redis"
	"github.com/astaxie/beego/logs"
)

func Init(s ...string) error {
	var err error
	err, nacos := nacos.ClientConfig()
	if err != nil {
		logs.Info(err, "nacos")
		return err
	}
	for _, s2 := range s {
		switch s2 {
		case "mysql":
			err = mysql.InitMysql(nacos)
		case "redis":
			redis.InitRedis()
		}
	}
	return err
}
