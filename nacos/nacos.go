package nacos

import (
	"github.com/AoShiJ/framework/config"
	"github.com/astaxie/beego/logs"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func ClientConfig() (error, string) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: config.N.Ip,
			Port:   uint64(config.N.Port),
		},
	}
	cc, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		logs.Info(err, "nacos")
		return err, ""
	}
	content, err := cc.GetConfig(vo.ConfigParam{
		DataId: config.N.DataId,
		Group:  config.N.Group})
	return nil, content
}
