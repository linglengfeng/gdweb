package main

import (
	"fmt"
	"webServer/request"
	"webServer/src/db"
	"webServer/src/gogrpc"
	"webServer/src/log"
	"webServer/src/sendgrid"
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
