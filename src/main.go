package main

import (
	"fmt"
	"gobasictinyurl/src/hashutility"
	"gobasictinyurl/src/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	srv := server.New(mux, ":3333")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		endpoint := r.URL.Query().Get("endpoint")
		w.Write([]byte(hashutility.GetNextHashSeq(endpoint)))
	})

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting the server...")
	}
}
