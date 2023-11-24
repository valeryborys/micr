package main

import (
	"micr/second/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.HandleBasicRequest)
	http.HandleFunc("/msg", api.SendKafkaMessage)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
