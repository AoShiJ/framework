package grpc

import (
	"fmt"
	"github.com/AoShiJ/framework/consul"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(toService, IP string, port int) (*grpc.ClientConn, error) {
	conn, err := consul.AgentHealthService(toService, IP, port)
	logs.Info(conn, 123)
	logs.Info(toService, 321)
	if err != nil {
		return nil, err
	}
	fmt.Println(conn)
	return grpc.Dial(conn, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
