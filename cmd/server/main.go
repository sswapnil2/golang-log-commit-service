package main

import (
	"log"

	"github.com/sswapnil2/proglog/internal/server"
)

func main(){

	server := server.NewHTTPServer(":8000")
	log.Fatal(server.ListenAndServe())
}
