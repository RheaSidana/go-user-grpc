package deferring

import "gRPC-go/server/initializer"

func DeferServer() {
	initializer.Server.Stop()
}