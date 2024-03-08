package grpc

import (
	"fmt"
	"github.com/AoShiJ/framework/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

func RegisterGrpc(port string, fc func(s *grpc.Server)) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	atoi, err := strconv.Atoi(port)
	if err != nil {
		return
	}

	consul.RegisterConsul(int64(atoi), "10.2.171.80", "sss")
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	fc(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
