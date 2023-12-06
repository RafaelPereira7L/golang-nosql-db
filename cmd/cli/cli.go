package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
	"github.com/rafaelpereira7l/golang-nosql-db/cmd/db"
)

func CreateNewCollection() {
	fmt.Println("Type the name of the collection to create: ")
	var name string
	fmt.Scanln(&name)

	db.CreateCollection(name)
}

func main() {
	Box := box.New(box.Config{Px: 10, Py: 5, Type: "Round", Color: "White", TitlePos: "Top"})
	Box.Println("MyDB", "NoSQL Database v1.0.0\n\n By Rafael Pereira\n\n Type a number to select an option:\n 1) Create a new collection\n 2) Add a new record\n 3) List all records\n 4) Get a record by ID\n 5) Delete a record\n 6) Exit")

	var option int
	fmt.Scanln(&option)

	switch option {
		case 1:
			CreateNewCollection()
		case 2:
			color.Red.Println("Not implemented yet")

		case 3:
			color.Red.Println("Not implemented yet")

		case 4:
			color.Red.Println("Not implemented yet")

		case 5:
			color.Red.Println("Not implemented yet")

		case 6:
			color.Yellow.Println("Bye!")

		default:
			color.Red.Println("Invalid option");
	}
}