package gogrpc

import (
	"context"
	"fmt"
	"web3Server/pkg/myutil"
	"web3Server/proto_go"
	"web3Server/src/db"
	"web3Server/src/mailer"
)

func (s *server) Say(ctx context.Context, req *proto_go.SayRequest) (*proto_go.SayResponse, error) {
	fmt.Println("svr Say request:", req)
	return &proto_go.SayResponse{Message: "my name is:" + req.Name}, nil
}

func (s *server) UserLogincode(ctx context.Context, req *proto_go.C2S_UserLogincode) (*proto_go.S2C_UserLogincode, error) {
	// fmt.Println("svr UserLogincode request:", req)
	account := req.Account
	code := myutil.RandomString(6)
	is, err := db.SetUserLoginCode(account, code)
	status := 1
	msg := ""
	if !is && err != nil {
		status = 0
		msg = err.Error()
	}
	err = mailer.SendLoginEmail(account, code)
	if err != nil {
		status = 0
		msg = err.Error()
	}
	fmt.Println("svr UserLogincode request:", req, "account:", account, "code:", code)
	// err1 := fmt.Errorf("sssss:%v", account)
	return &proto_go.S2C_UserLogincode{Status: int32(status), Message: msg}, nil
}

func (s *server) UserLoginauth(ctx context.Context, req *proto_go.C2S_UserLoginauth) (*proto_go.S2C_UserLoginauth, error) {
	fmt.Println("svr UserLoginauth request:", req)
	return &proto_go.S2C_UserLoginauth{}, nil
}

func (s *server) CommMsg(ctx context.Context, req *proto_go.C2S_Map) (*proto_go.S2C_Map, error) {
	fmt.Println("svr CommMsg request:", req)
	return &proto_go.S2C_Map{Map: req.Map}, nil
}
