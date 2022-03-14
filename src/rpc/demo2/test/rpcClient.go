package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc/demo2/hello"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ hello.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(hello.HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
