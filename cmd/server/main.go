package main

import (
	"fmt"
	"log"

	"github.com/sswapnil2/proglog/internal/server"
)

const (
	ServerPort string = ":8000"
)

func main() {
	fmt.Println(fmt.Sprintf("Starting server: %s", ServerPort))
	srv := server.NewHTTPServer(ServerPort)
	log.Fatal(srv.ListenAndServe())
}
