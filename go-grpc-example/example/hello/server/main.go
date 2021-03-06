package main

import (
    "net"
    pb "github.com/xingcuntian/go_test/go-grpc-example/example/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
)

const (
    //Address gRPC服务地址
    Address = "127.0.0.1:50052"
)
//定义helloService并实现约定的接口
type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    resp := new(pb.HelloReply)
    resp.Messge = "Hello " + in.Name + "."
    return resp, nil
}

func main(){
    listen,err:=net.Listen("tcp",Address)
    if err!=nil{
        grpclog.Fatalf("failed to listen: %v",err)
    }
    //实例化gRPC Server
    s:=grpc.NewServer()
    //注册HelloService
    pb.RegisterHelloServer(s,HelloService)    
    grpclog.Println("Listen on " + Address)
    s.Serve(listen)
}