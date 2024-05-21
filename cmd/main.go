package main

import (
	"fmt"
	"log"

	"github.com/JMustang/warehouse/cmd/api"
	"github.com/JMustang/warehouse/db"
)

func main() {
	db := db.GetDB()
	defer db.Close()

	fmt.Println("Connection successful!")

	server := api.NewAPIServer(":8000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
