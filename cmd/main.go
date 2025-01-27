package main

import (
	"log"
	"net/http"

	"github.com/lutlevon/expense-mate/config"
	"github.com/lutlevon/expense-mate/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := routes.Routes()
	log.Println("server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
