package services

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"my-web-app/protocol/gen-go/user"
)

type HelloServiceHandler struct {
}

func (h *HelloServiceHandler) SayHello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

//func InitServer() {
//	transport, err := thrift.NewTServerSocket("localhost:9091")
//	if err != nil {
//		fmt.Println("Error starting server:", err)
//		return
//	}
//	handler := &HelloServiceHandler{}
//	processer := tutoral.NewHelloServiceProcessor(handler)
//	server := thrift.NewTSimpleServer4(processer, transport, thrift.NewTTransportFactory(), thrift.NewTBinaryProtocolFactoryDefault())
//	fmt.Println("Starting the server...", transport)
//	if err = server.Serve(); err != nil {
//		fmt.Println("Error starting server:", err)
//	}
//}

func InitUserService() {
	transport, err := thrift.NewTServerSocket("localhost:9091")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	handler := &UserServiceHandler{}
	processer := user.NewUserServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processer, transport, thrift.NewTTransportFactory(), thrift.NewTBinaryProtocolFactoryDefault())
	fmt.Println("Starting the server...", transport)
	if err = server.Serve(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

type UserServiceHandler struct {
}

func (h *UserServiceHandler) GetUser(ctx context.Context, userId int32) (*user.User, error) {
	fmt.Println("userId=", userId)
	return &user.User{
		ID:   userId,
		Name: "xiaoming",
	}, nil
}

func (h *UserServiceHandler) UpdateUser(ctx context.Context, user *user.User) error {
	fmt.Println("update user...")
	return nil
}

func (h *UserServiceHandler) GetAllUsers(ctx context.Context) ([]*user.User, error) {
	var userList []*user.User
	userList = append(userList, &user.User{
		ID:   0,
		Name: "xiaoming",
	})
	return userList, nil
}
