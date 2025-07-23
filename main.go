package main

import (
	"log"
	"net/http"

	"api/routers"
)

func main() {
	router := routers.NewRouter()

	log.Println("Server started on :3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
