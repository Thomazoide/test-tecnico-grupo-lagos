package main

import (
	"fmt"

	"github.com/test-tecnico-grupo-lagos/backend/config"
	"github.com/test-tecnico-grupo-lagos/backend/db"
)

func main() {
	db.ConnectDB()
	server := config.NewApiServer(":8080")
	server.RunServer()
	fmt.Println("Server on port 8080")
}
