package app

import (
	"github.com/astaxie/beego/logs"
	"week2/rpc/work/mysql"
	"week2/rpc/work/nacos"
	"week2/rpc/work/redis"
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
