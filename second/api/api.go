package api

import (
	"fmt"
	"micr/second/service"
	"net/http"
)

func HandleBasicRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my website!")
}

func SendKafkaMessage(w http.ResponseWriter, r *http.Request) {
	service.SendMessage()
}
