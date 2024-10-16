package client

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"my-web-app/protocol/gen-go/tutoral"
)

func GetClient() *tutoral.HelloServiceClient {
	addr := ":9090"
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := tutoral.NewHelloServiceClient(thrift.NewTStandardClient(iprot, oprot))
	return client
}

func HelloClient(ctx context.Context) string {
	client := GetClient()
	rep, err := client.SayHello(ctx, "test")
	fmt.Println(rep, err)
	if err != nil {
		fmt.Println("thrift err: %v", err)
	}
	return rep
}
