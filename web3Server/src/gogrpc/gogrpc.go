package gogrpc

import (
	"fmt"
	"net"
	"web3Server/config"
	"web3Server/pkg/ggrpc"
	"web3Server/proto_go"

	"google.golang.org/grpc"
)

var (
	Svr            *grpc.Server
	SvrNetListener net.Listener
	CliConn        *grpc.ClientConn
	Cli            client
)

type (
	server struct {
		proto_go.UnimplementedHelloServer
		proto_go.UnimplementedGogrpcServer
	}
	client struct {
		HelloClient  proto_go.HelloClient
		GogrpcClient proto_go.GogrpcClient
	}
)

func registerSvr(s grpc.ServiceRegistrar) {
	proto_go.RegisterHelloServer(s, &server{})
	proto_go.RegisterGogrpcServer(s, &server{})
}

func newCli(conn *grpc.ClientConn) {
	Cli.HelloClient = proto_go.NewHelloClient(conn)
	Cli.GogrpcClient = proto_go.NewGogrpcClient(conn)
}

func Start() {
	go startSvr()
	go startCli()
}

func Stop() {
	stopSvr()
	stopCli()
}

func startSvr() {
	grpcport := config.Config.GetString("grpc.svr.port")
	if grpcport == "" {
		return
	}
	svr, netListener, err := ggrpc.StartSvr(grpcport)
	if err != nil {
		fmt.Println("gogrpc startSvr failed, err:", err)
	}
	registerSvr(svr)
	Svr = svr
	SvrNetListener = netListener
	fmt.Println("gogrpc startSvr successed...", "port:", grpcport)
	ggrpc.ListenSvr()
}

func stopSvr() {
	grpcport := config.Config.GetString("grpc.svr.port")
	if grpcport == "" {
		return
	}
	ggrpc.StopSvr()
}

var HelloClient proto_go.HelloClient
var GogrpcClient proto_go.GogrpcClient

func startCli() {
	grpc_ip := config.Config.GetString("grpc.cli.grpc_ip")
	grpc_port := config.Config.GetString("grpc.cli.grpc_port")
	if grpc_ip == "" || grpc_port == "" {
		return
	}
	conn, err := ggrpc.StartCli(grpc_ip, grpc_port)
	if err != nil {
		fmt.Println("gogrpc startCli failed, err:", err)
	}
	newCli(conn)
	CliConn = conn
	fmt.Println("gogrpc startCli successed..., grpc_ip:", grpc_ip, ", grpc_port:", grpc_port)
}

func stopCli() {
	CliConn.Close()
}
