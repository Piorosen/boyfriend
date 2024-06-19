package main

import (
	"C"
	"fmt"
)

//export helloworld
func helloworld() {
	fmt.Printf("hello world!")
}

//export helloworld1
func helloworld1(data int) {
	fmt.Printf("hello world! %d", data)
}

//export helloworld2
func helloworld2(data int) {
	fmt.Printf("hello world! %d", data)
}

func main() {}
