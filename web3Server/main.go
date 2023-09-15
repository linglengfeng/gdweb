package main

import (
	"fmt"
	"web3Server/request"
	"web3Server/src/db"
	"web3Server/src/gogrpc"
	"web3Server/src/log"
	"web3Server/src/sendgrid"
)

func main() {
	log.Start()
	fmt.Println("log start successed...")
	db.Start()
	fmt.Println("db start successed...")
	sendgrid.Start()
	fmt.Println("sendgrid start successed...")
	gogrpc.Start()
	request.Start()
}
