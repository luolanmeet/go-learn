package hello

import (
	"fmt"
	"log"
	"net"
)

type HelloService struct {
	Conn    net.Conn
	isLogIn bool
}

func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogIn = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogIn {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ", from " + p.Conn.RemoteAddr().String()
	return nil
}
