package main

import (
	"context"
	"crypto/tls"
	grpctesting "github.com/soichisumi-sandbox/grpc-bidistream-lb/proto"
	"github.com/soichisumi/go-util/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func main(){
	endpoint := "34.120.219.175:443"
	cownn, err := grpc.Dial(
		endpoint,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: 150 * time.Second, PermitWithoutStream: true}),
	)
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	c := grpctesting.NewEchoServiceClient(conn)

	// unary rpc
	//res, err := c.Echo(context.Background(), &grpctesting.EchoRequest{Message: "yo"})
	//if err != nil {
	//	logger.Fatal("", zap.Error(err))
	//}
	//logger.Info("", zap.String("msg", res.GetMessage()))

	// client stream
	//stream, err := c.EchoClientStream(context.Background())
	//if err != nil {
	//	logger.Fatal("", zap.Error(err))
	//}
	//for i:=0; i<5; i++ {
	//	if err := stream.Send(&grpctesting.EchoRequest{Message: strconv.Itoa(rand.Int())}); err != nil {
	//		logger.Error("", zap.Error(err))
	//	}
	//}
	//recv, err := stream.CloseAndRecv()
	//if err != nil {
	//	logger.Error("", zap.Error(err))
	//}
	//logger.Info("", zap.String("msg", recv.GetMessage()))

	// server stream
	//stream, err := c.EchoServerStream(context.Background(), &grpctesting.EchoRequest{
	//	Message: "wiiiiiiii",
	//})
	//if err != nil {
	//	logger.Fatal("", zap.Error(err))
	//}
	//for {
	//	recv, err := stream.Recv()
	//	if err == io.EOF{
	//		break
	//	}
	//	if err != nil {
	//		logger.Fatal("", zap.Error(err))
	//	}
	//	logger.Info("", zap.String("msg", recv.GetMessage()))
	//}

	// bidi stream
	stream, err := c.EchoBidiStream(context.Background())
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	waitChan := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitChan)
				return
			}
			if err != nil {
				logger.Fatal("", zap.Error(err))
			}
			logger.Info("", zap.String("msg", in.GetMessage()))
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	count := 0
	for {
		select {
		case <- ticker.C:
			if err := stream.Send(&grpctesting.EchoRequest{Message: strconv.Itoa(rand.Int())}); err != nil {
				logger.Fatal("", zap.Error(err))
			}
			count++
		}
		if count > 60 {
			break
		}
	}
	stream.CloseSend()
	<-waitChan
}
