package hello

import "net/rpc"

// 抽象出一个接口，存放约定的常量
const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
