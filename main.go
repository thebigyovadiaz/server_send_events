package main

import (
	"github.com/thebigyovadiaz/server_send_events/handlers"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	handlers.InitRoutes(r)
	log.Println("Starting server >>> localhost:8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}
