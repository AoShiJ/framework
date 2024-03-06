package nacos

import (
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
			IpAddr: "192.168.127.9",
			Port:   8848,
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
		DataId: "sss",
		Group:  "DEFAULT_GROUP"})
	return nil, content
}
