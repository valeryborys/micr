package api

import (
	"micr/forth/service"
	"net/http"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
)

func RegisterUserRouters(r *mux.Router) {
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/", Welcome).Methods("GET")
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsoniter.NewEncoder(w).Encode("Welcome message, application is running")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ret, err := service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), 404) //TODO proper error handling
	}

	jsoniter.NewEncoder(w).Encode(ret)
}
