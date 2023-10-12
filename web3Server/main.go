package main

import (
	"fmt"
	"web3Server/request"
	"web3Server/src/db"
	"web3Server/src/gogrpc"
	"web3Server/src/httpcli"
	"web3Server/src/log"
	"web3Server/src/mailer"
)

func main() {
	log.Start()
	fmt.Println("log start successed...")
	db.Start()
	fmt.Println("db start successed...")
	mailer.Start()
	fmt.Println("mailer start successed...")
	gogrpc.Start()
	fmt.Println("gogrpc start successed...")
	httpcli.Start()
	fmt.Println("httpcli start successed...")
	request.Start()
}
