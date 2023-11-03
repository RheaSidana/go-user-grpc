package main

import (
	deferring "gRPC-go/server/defer"
	"gRPC-go/server/initializer"
	users "gRPC-go/server/module/user"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.CreateNewGrpcServer()
}

func main() {
	defer deferring.DeferAllInit()
	
	users.RegisterUserServer()
	users.ListenUserServer()
}
