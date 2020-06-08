package main

import (
	"context"
	grpctesting "github.com/soichisumi-sandbox/grpc-bidistream-lb/proto"
	"github.com/soichisumi/go-util/logger"
	"go.uber.org/zap"
	"io"
	"math/rand"
	"strings"
)

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

type EchoServer struct{}

func (EchoServer) Echo(ctx context.Context, req *grpctesting.EchoRequest) (*grpctesting.EchoResponse, error) {
	logger.Info("", zap.Any("req", req))
	return &grpctesting.EchoResponse{
		Message: req.Message,
	}, nil
}

func (EchoServer) Empty(ctx context.Context, req *grpctesting.EmptyRequest) (*grpctesting.EmptyResponse, error) {
	logger.Info("", zap.Any("req", req))
	return &grpctesting.EmptyResponse{}, nil
}

func (s EchoServer) EchoClientStream(stream grpctesting.EchoService_EchoClientStreamServer) error {
	msgs := make([]string, 0, 0)
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&grpctesting.EchoResponse{Message: strings.Join(msgs, ",")})
		}
		if err != nil {
			return err
		}
		logger.Info("msg", zap.String("msg", recv.GetMessage()))
		msgs = append(msgs, recv.GetMessage())
	}
}

func (s EchoServer) EchoServerStream(request *grpctesting.EchoRequest, stream grpctesting.EchoService_EchoServerStreamServer) error {
	logger.Info("EchoServerStream", zap.String("msg", request.GetMessage()))
	for i := 0; i < rand.Intn(5); i++ {
		if err := stream.Send(&grpctesting.EchoResponse{Message: request.GetMessage()}); err != nil {
			return err
		}
	}
	return nil
}

func (s EchoServer) EchoBidiStream(stream grpctesting.EchoService_EchoBidiStreamServer) error {
	msgs := make([]string, 0, 0)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		logger.Info("EchoBidiStream", zap.String("msg", in.GetMessage()))

		msgs = append(msgs, in.GetMessage())

		for i := 0; i < rand.Intn(5); i++ {
			if err := stream.Send(&grpctesting.EchoResponse{Message: strings.Join(msgs, ", ")}); err != nil {
				return err
			}
		}
	}
}
