package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)
import "google.golang.org/grpc"

func main() {
	fmt.Println("hello")
	conn := grpc.ClientConn{}
	buffer := proto.Buffer{}
	fmt.Println(conn)
	fmt.Println(buffer)
}
