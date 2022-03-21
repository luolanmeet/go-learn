package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	//err = client.Call("HelloService.Login", "root", &reply)
	err = client.Call("HelloService.Login", "user:password", &reply)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Call("HelloService.Hello", "root", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
