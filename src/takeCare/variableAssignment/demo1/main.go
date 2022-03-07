package main

import "fmt"

func main() {

	x := []int{1, 2, 3}
	i := 0
	i, x[i] = 1, 2 // set i = 1, x[0]=2
	fmt.Println(i, x)

	x = []int{1, 2, 3}
	i = 0
	x[i], i = 2, x[i] // set tmp=x[0], x[0]=2, i=tmp ==> i=1
	fmt.Println(i, x)

	x[0], x[0] = 1, 2 // set x[0]=1, then x[0]=2
	fmt.Println(x[0])
}
