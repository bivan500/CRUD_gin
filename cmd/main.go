package main

import (
	bookList "CRUD_GIN"
	"log"
)

func main() {
	srv := new(bookList.Server)
	err := srv.Run("8080")
	if err != nil {
		log.Fatal("error in server start: %v", err.Error())
	}
}
