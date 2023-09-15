package gogrpc

import (
	"context"
	"fmt"
	"web3Server/proto_go"

	"google.golang.org/grpc"
)

func (*client) Say(ctx context.Context, in *proto_go.SayRequest, opts ...grpc.CallOption) (*proto_go.SayResponse, error) {
	fmt.Println("cli Say request:", in)
	return &proto_go.SayResponse{Message: in.Name}, nil
}

func (c *client) UserLogincode(ctx context.Context, in *proto_go.C2S_UserLogincode, opts ...grpc.CallOption) (*proto_go.S2C_UserLogincode, error) {
	fmt.Println("cli UserLogincode request:", in)
	return &proto_go.S2C_UserLogincode{}, nil
}

func (c *client) UserLoginauth(ctx context.Context, req *proto_go.C2S_UserLoginauth, opts ...grpc.CallOption) (*proto_go.S2C_UserLoginauth, error) {
	fmt.Println("cli UserLoginauth request:", req)
	return &proto_go.S2C_UserLoginauth{}, nil
}

func (c *client) CommMsg(ctx context.Context, req *proto_go.C2S_Map, opts ...grpc.CallOption) (*proto_go.S2C_Map, error) {
	fmt.Println("cli CommMsg request:", req)
	return &proto_go.S2C_Map{Map: req.Map}, nil
}
