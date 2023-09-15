package gogrpc

import (
	"context"
	"fmt"
	"webServer/proto_go"
)

func (s *server) Say(ctx context.Context, req *proto_go.SayRequest) (*proto_go.SayResponse, error) {
	fmt.Println("svr Say request:", req)
	return &proto_go.SayResponse{Message: "my name is:" + req.Name}, nil
}

func (s *server) UserLogincode(ctx context.Context, req *proto_go.C2S_UserLogincode) (*proto_go.S2C_UserLogincode, error) {
	fmt.Println("svr UserLogincode request:", req)
	return &proto_go.S2C_UserLogincode{}, nil
}

func (s *server) UserLoginauth(ctx context.Context, req *proto_go.C2S_UserLoginauth) (*proto_go.S2C_UserLoginauth, error) {
	fmt.Println("svr UserLoginauth request:", req)
	return &proto_go.S2C_UserLoginauth{}, nil
}

func (s *server) CommMsg(ctx context.Context, req *proto_go.C2S_Map) (*proto_go.S2C_Map, error) {
	fmt.Println("svr CommMsg request:", req)
	return &proto_go.S2C_Map{Map: req.Map}, nil
}
