package auth

import (
	"encoding/json"
	"gobasictinyurl/src/models"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	user.HashPassword(user.Password)
}

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", RegisterUser)
}
