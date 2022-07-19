package main

import "fmt"

type Comment struct {
	Author string
	Body   string
	Slug   string
	ID     int
}

func main() {
	cmt := new(Comment)
	fmt.Printf("%+v\n", cmt)
}

// go run comment.go
