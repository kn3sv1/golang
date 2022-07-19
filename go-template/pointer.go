package main

import "fmt"

func main() {
	var a int = 6
	var age *int = &a

	fmt.Println(&a)
	fmt.Println(age)
	
	//address of this variable
	fmt.Println(&age)
}

// https://tutorialedge.net/golang/go-pointers-tutorial/

// go run pointer.go
