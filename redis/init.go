package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"week2/rpc/work/mysql"
)

var Red *redis.Client

func InitRedis() {
	Red = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%v:%v", mysql.T1.Redis.Ip, mysql.T1.Redis.Port)})
}
