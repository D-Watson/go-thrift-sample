package client

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"my-web-app/protocol/gen-go/user"
)

func GetClient() *user.UserServiceClient {
	addr := ":9091"
	//声明一个 thrift.TTransport 类型的变量，用于管理与 Thrift 服务的连接
	var transport thrift.TTransport
	var err error
	//使用 thrift.NewTSocket 创建一个新的套接字（socket）传输，连接到指定的地址 addr,该方法返回一个 TTransport 实例和可能的错误。
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	//声明一个 thrift.TProtocolFactory 类型的变量，并利用 thrift.NewTBinaryProtocolFactory 创建一个新的二进制协议工厂,用于序列化和反序列化数据
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//声明一个 thrift.TTransportFactory 类型的变量，并创建一个新的传输工厂。这里使用的是默认的无缓冲传输工厂。
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()
	//使用传输工厂的 GetTransport 方法来获取一个新的传输对象,这里传入之前创建的 transport
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}
	//尝试打开传输连接, 若打开失败，则打印错误信息
	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}
	//iprot输入协议
	iprot := protocolFactory.GetProtocol(transport)
	//oprot输出协议
	oprot := protocolFactory.GetProtocol(transport)
	//使用输入和输出协议创建一个新的 UserServiceClient 实例。thrift.NewTStandardClient 用于创建一个标准的 Thrift 客户端
	client := user.NewUserServiceClient(thrift.NewTStandardClient(iprot, oprot))
	return client
}

func HelloClient(ctx context.Context) *user.User {
	client := GetClient()
	rep, err := client.GetUser(ctx, 1)
	fmt.Println(rep, err)
	if err != nil {
		fmt.Println("thrift err:", err)
	}
	return rep
}
