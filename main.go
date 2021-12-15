package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vairarchi/jwt-go-example/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/refresh", handlers.Refresh)

	fmt.Println("Server started on port :8081")
	log.Panic(http.ListenAndServe(":8081", nil))
}
