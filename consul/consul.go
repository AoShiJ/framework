package consul

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"net"
)

const Consul_Key = "consul_index"

type Con struct {
	Consul struct {
		Ip   string `json:"Ip"`
		Port string `json:"Port"`
	} `json:"Consul"`
}

//	func getConfig(servername string) (*Con, error) {
//		s := new(Con)
//		config, err := GetConfig(servername, "DEFAULT_GROUP")
//		if err != nil {
//			return nil, err
//		}
//		json.Unmarshal([]byte(config), &s)
//		return s, err
//	}
func GetIp() (ip []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, addr := range addrs {
		ipNet, isVailIpNet := addr.(*net.IPNet)
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = append(ip, ipNet.IP.String())
			}
		}

	}
	return ip
}

func RegisterConsul(servername string, IP string, port int) error {
	//config, err := getConfig(servername)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(config.Consul.Ip, config.Consul.Port, 2342423423424324)

	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%v:%v", IP, "8500"),
	})
	if err != nil {
		return err
	}
	ip := GetIp()
	logs.Info(client)
	logs.Info(ip[0])
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    servername,
		Tags:    []string{"Grpc"},
		Port:    port,
		Address: ip[0],
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			GRPC:                           fmt.Sprintf("%v:%v", ip[0], port),
			DeregisterCriticalServiceAfter: "10s",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func AgentHealthService(servername string, IP string, port int) (string, error) {
	//config, err := getConfig(servername)
	//if err != nil {
	//	return "", err
	//}
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%v:%v", IP, port),
	})
	if err != nil {
		return "", err
	}
	name, i, err := client.Agent().AgentHealthServiceByName(servername)
	if err != nil {
		return "", err
	}
	if name != "passing" {
		return "", fmt.Errorf("is not health service")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v:%v", i[0].Service.Address, i[0].Service.Port), nil
}
