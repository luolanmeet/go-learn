package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

/**
 * 结构的方法，不是包在结构体定义中，而是和结构体同级
 * 把接收的结构放在方法前的括号中
 * 和普通的方法只有形式上的区别
 */
func (node *treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

/**
 * Go是值传递！这里接收的是一个struct，修改value不会影响原struct
 */
func (node treeNode) setValue1(value int) {
	node.value = value
}

/**
 * 测试下深拷贝还是浅拷贝
 * 结果是浅拷贝，原node的right结构体的value会发生变化
 */
func (node treeNode) setRightValue(value int) {
	node.right.value = value
}

/**
 * 要想控制struct的创建，可以对外提供工厂函数
 */
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func define() {

	var root treeNode
	// 这里不用&，因为创建了一个struct
	root = treeNode{value: 1}

	// 这里需要 &，因为是要给一个指针去赋值
	// 注意到直接用 struct 的. 可以访问到struct里的成员
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	// new内置函数，可以直接返回的struct的地址
	// 这里left是指针，注意到同样直接用.可以访问到指定所指的struct的成员
	root.left.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{1, nil, &root},
	}

	fmt.Println(nodes)

	root.print()
	root.setValue1(10)
	root.print()
	root.setValue(11)
	root.print()

	fmt.Println("test deep copy")
	root.right.print()
	root.setRightValue(10)
	root.right.print()
}

func main() {

	define()

}
