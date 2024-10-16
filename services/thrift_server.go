package services

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"my-web-app/protocol/gen-go/tutoral"
)

type HelloServiceHandler struct {
}

func (h *HelloServiceHandler) SayHello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

func InitServer() {
	transport, err := thrift.NewTServerSocket("localhost:9090")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	handler := &HelloServiceHandler{}
	processer := tutoral.NewHelloServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processer, transport, thrift.NewTTransportFactory(), thrift.NewTBinaryProtocolFactoryDefault())
	fmt.Println("Starting the server...")
	if err = server.Serve(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
