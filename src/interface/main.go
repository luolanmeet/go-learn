package main

import (
	"fmt"
	"interface/mock"
	"interface/real"
	"time"
)

type Retriever interface {
	// interface 里，方法不用写 func
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {

	// 使用方（main），定义了接口
	// 只要方法声明对得上，就可以做为接口的实现方，不需要真的和定义的接口有关联关系
	// 只有方法而已，注意，是满足接口的所有方法。
	// 如何使用方的接口变了，在提供方还没被用到的情况下，提供方是不知道的.. 因为没用的，编译器也无法检查出来

	mock := mock.RetrieverB{"hiiiiiiii"}
	mock.Post("222")
	fmt.Println(download(mock))

	var real Retriever = &real.RetrieverA{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	// fmt.Println(download(real))

	// mock real 两个接口变量里的内容，是实现者的类型，还有实现者的值或指针，
	fmt.Printf("%T %v\n", mock, mock)
	fmt.Printf("%T %v\n", real, real)

}
