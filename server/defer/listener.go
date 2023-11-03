package deferring

import "gRPC-go/server/initializer"

func DeferListener() {
	initializer.Lis.Close()
}

