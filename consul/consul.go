package consul

import (
	"context"
	"fmt"
	"github.com/AoShiJ/framework/redis"
	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"strconv"
	"time"
)

const CONSUL_KEY = "consul:node:index"

func RegisterConsul(port int64, address, name string) error {

	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}

	err = c.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    name,
		Tags:    []string{"GRPC"},
		Port:    int(port),
		Address: address,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",                                //间隔时常
			Timeout:                        "5s",                                //退出
			GRPC:                           fmt.Sprintf("%v:%v", address, port), //
			DeregisterCriticalServiceAfter: "30s",                               //注销
		},
	})
	if err != nil {
		return err
	}
	return nil
}
func getIndex(ctx context.Context, serviceName string, indexLen int) (int, error) {
	exist, err := redis.ExistKey(ctx, serviceName, CONSUL_KEY)
	if err != nil {
		return 0, err
	}

	if exist {
		indexStr, err := redis.GetByKey(ctx, serviceName, CONSUL_KEY)
		if err != nil {
			return 0, err
		}
		index, err := strconv.Atoi(indexStr)
		newIndex := index + 1

		if newIndex >= indexLen {
			newIndex = 0
		}
		err = redis.SetKey(ctx, serviceName, CONSUL_KEY, newIndex, time.Duration(0))
		if err != nil {
			return 0, err
		}

		return index, nil
	}

	err = redis.SetKey(ctx, serviceName, CONSUL_KEY, 0, time.Duration(0))
	if err != nil {
		return 0, err
	}
	return 0, nil
}
func AgentHealthService(ctx context.Context, serviceName string) (string, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return "", err
	}
	sr, infos, err := client.Agent().AgentHealthServiceByName(serviceName)
	logs.Info(sr)
	logs.Info(infos, 123123)
	if err != nil {
		return "", err
	}
	if sr != "passing" {
		return "", fmt.Errorf("is not have health service")
	}

	index, err := getIndex(ctx, serviceName, len(infos))
	if err != nil {
		return "", err
	}
	logs.Info(index, "-------------index")
	logs.Info(fmt.Sprintf("%v:%v", infos[index].Service.Address, infos[index].Service.Port), "---------infos")
	return fmt.Sprintf("%v:%v", infos[index].Service.Address, infos[index].Service.Port), nil
}
