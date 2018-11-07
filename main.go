package main

import (
	"log"
	"net/http"

	"orchestra/api"
)

func main() {
	err := http.ListenAndServe(":3000", api.Handlers())
	//err := http.ListenAndServeTLS(":3000", "certificate/localhost.pem", "certificate/localhost.key", api.Handlers())

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("Server : localhost:3000")

}
