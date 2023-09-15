package ggrpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Svr *grpc.Server
var SvrNetListener net.Listener

func StartSvr(grpcport string) (*grpc.Server, net.Listener, error) {
	s := grpc.NewServer()
	listen, err := net.Listen("tcp", ":"+grpcport)
	Svr = s
	SvrNetListener = listen
	if err != nil {
		return s, listen, err
	}
	// defer stop()
	return s, listen, nil
}

func ListenSvr() error {
	return Svr.Serve(SvrNetListener)
}

func StopSvr() {
	Svr.Stop()
	SvrNetListener.Close()
}

func StartCli(ip, port string) (*grpc.ClientConn, error) {
	serviceHost := ip + ":" + port
	return grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
