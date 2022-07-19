package main

import (
	"fmt"
	"log"

	wmenu "github.com/dixonwille/wmenu/v5"
)

func handleFunc(opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		fmt.Println("Adding a new Person")
	case 1:
		fmt.Println("Finding a Person")
	case 2:
		fmt.Println("Update a Person's information")
	case 3:
		fmt.Println("Deleting a person by ID")
	case 4:
		fmt.Println("Quitting application")
	}
}

func main() {

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(opts); return nil })

	menu.Option("Add a new Person", 0, true, nil)
	menu.Option("Find a Person", 1, false, nil)
	menu.Option("Update a Person's information", 2, false, nil)
	menu.Option("Delete a person by ID", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

// go run main.go

// https://www.allhandsontech.com/programming/golang/how-to-use-sqlite-with-go/

// go mod init sqllite3
// go get github.com/dixonwille/wmenu/v5
