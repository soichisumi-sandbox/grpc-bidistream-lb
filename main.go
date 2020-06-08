package main

import (
	"flag"
	"fmt"
	grpctesting "github.com/soichisumi-sandbox/grpc-bidistream-lb/proto"
	"github.com/soichisumi/go-util/logger"
	"github.com/soichisumi/grpc-echo-server/pkg/health"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

const (
	defaultPort = 8080
)

func main() {
	var port = flag.Int("p", defaultPort, "port number for listening")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	//creds, err := credentials.NewServerTLSFromFile("./certs/cert.pem", "./certs/privkey.pem")
	//if err != nil {
	//	logger.Fatal(err.Error(), zap.Error(err))
	//}
	//server := grpc.NewServer(grpc.Creds(creds))

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{Time: 150 * time.Second}))
	grpctesting.RegisterEchoServiceServer(server, NewEchoServer())
	grpc_health_v1.RegisterHealthServer(server, health.NewHealthServer())
	reflection.Register(server)

	logger.Info("", zap.Int("port", *port))
	if err := server.Serve(lis); err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
}