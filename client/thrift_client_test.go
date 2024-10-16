package client

import (
	"context"
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	client := GetClient()
	rep, err := client.GetUser(context.Background(), 1)
	fmt.Println(rep, err)
	if err != nil {
		fmt.Println("thrift err:", err)
	}
}
