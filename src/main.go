// KV server entry point.
package main

import (
	"database"
	"fmt"
	"server"
)

func main() {
	fmt.Println("Starting database")

	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	fmt.Println("Starting server")

	s := server.MakeServer(db)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err.Error())
	}
}
