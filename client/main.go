package main

import (
	deferring "gRPC-go/client/defer"
	"gRPC-go/client/initializer"
	users "gRPC-go/client/module/user"
	"log"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.CreateNewGrpcClient()
}

func main() {
	defer deferring.DeferConn()

	users.NewUserClient()

	log.Println("*** *** GetUserDetails *** ***")

	userRequest := users.GenerateRequest(1)
	users.CallGetUserDetails(userRequest)

	userRequest = users.GenerateRequest(3)
	users.CallGetUserDetails(userRequest)

	log.Println("*** *** CallGetUserDetailsByIDs *** ***")
	userIds := []int32{1,3,4}
	users.CallGetUserDetailsByIDs(userIds)
}
