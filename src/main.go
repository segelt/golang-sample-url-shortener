package main

import (
	"fmt"
	"gobasictinyurl/src/hashutility"
	"gobasictinyurl/src/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	hashutility.SetupRoutes(mux)

	srv := server.New(mux, ":3333")
	err := srv.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting the server...")
	}
}
