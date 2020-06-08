# grpc-bidistream-lb

A sample repo for gRPC bidirectional stream with GCP External HTTP/S Load Balancer

## Server specification

```proto
syntax = "proto3";

package grpctesting;

message EmptyRequest {}
message EmptyResponse {}

message EchoRequest {
    string message = 1;
}

message EchoResponse {
    string message = 1;
}

service EchoService {
    rpc Echo(EchoRequest) returns (EchoResponse);
    rpc Empty(EmptyRequest) returns (EmptyResponse);
    rpc EchoClientStream(stream EchoRequest) returns (EchoResponse);
    rpc EchoServerStream(EchoRequest) returns (stream EchoResponse);
    rpc EchoBidiStream(stream EchoRequest) returns (stream EchoResponse);
}
```