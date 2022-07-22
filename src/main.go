package main

import (
	"fmt"
	"gobasictinyurl/src/auth"
	"gobasictinyurl/src/hashutility"
	"gobasictinyurl/src/persistence"
	"gobasictinyurl/src/server"
	"net/http"
)

func main() {
	persistence.Connect("postgres://sampleuser:12345@localHost:5432/golangtinyurldb")
	persistence.Migrate()

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
