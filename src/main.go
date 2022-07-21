package main

import (
	"fmt"
	"gobasictinyurl/src/auth"
	"gobasictinyurl/src/hashutility"
	"gobasictinyurl/src/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	hashpage := hashutility.NewHashPage()
	hashpage.SetupRoutes(mux)
	auth.SetupRoutes(mux)

	srv := server.New(mux, ":3333")
	err := srv.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting the server...")
	}
}
