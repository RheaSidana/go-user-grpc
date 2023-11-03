<h1>User-gRPC-Go</h1>

<h3> Steps to run the application </h3>
<h4>  1. Clone the repo </h4>
<h4>  2. run the command: go mod tidy </h4>
<h4>  3. Install the gRPC Go plugin </h4>

```
From:  https://grpc.io/docs/languages/go/quickstart/
```

<h4>  4. Protocol Buffer Compiler Installation </h4>

```
From: https://grpc.io/docs/protoc-installation/
```

<h4>Optional: delete ".\proto\users.pb.go" and ".\proto\users_grpc.pb.go" files
<br/>
Then run the following command:
</h4>

```
protoc --go_out=. --go-grpc_out=. proto/users.proto
```

<h4>  5. Change directory to Server in one terminal: </h4>

```
cd .\server
go run .
```

<h4>  6. Change directory to Client in one terminal: </h4>

```
cd .\client
go run .
```

<h4>  6. See the output on client terminal</h4>

