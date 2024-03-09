package grpc

import (
	"fmt"
	"github.com/AoShiJ/framework/consul"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RegisterGrpc(IP, serverName string, port int, fc func(s *grpc.Server)) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	err = consul.RegisterConsul(serverName, IP, port)
	if err != nil {
		logs.Info(err, "------consul")
		return
	}
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	fc(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
