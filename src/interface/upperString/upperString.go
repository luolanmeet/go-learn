package main

import (
	"fmt"
	"os"
	"strings"
)

type UpperString string

func (us UpperString) String() string {
	return strings.ToUpper(string(us))
}

func main() {
	fmt.Fprintln(os.Stdout, UpperString("hello world"))
}
