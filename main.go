package main

import (
	"log"
	"net/http"

	"github.com/tknkaa/wordle/api"
)

func main() {
	router := api.SetupRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
