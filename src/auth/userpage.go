package auth

import (
	"encoding/json"
	"fmt"
	"gobasictinyurl/src/models"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	user.HashPassword(user.Password)

	err = AddUserToStore(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(fmt.Sprintf("User registration complete. username: %s .", user.Name)))
	}

}

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request TokenRequest
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}

	user, err := GetUserFromStorage(request.Email)

	if err != nil {
		panic(err)
	}

	credentialError := user.CheckPassword(request.Password)

	if credentialError != nil {
		panic(credentialError)
	}

	tokenString, err := GenerateJWT(user.Email, user.Username)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", RegisterUser)
	mux.HandleFunc("/login", GenerateToken)
}
