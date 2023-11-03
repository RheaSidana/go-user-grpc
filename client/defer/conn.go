package deferring

import "gRPC-go/client/initializer"

func DeferConn() {
	initializer.Conn.Close()
}