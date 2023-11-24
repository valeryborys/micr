package main

import (
	"context"
	"log"
	"micr/forth/api"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := 8081 //TODO variable
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		log.Println("signal caught. shutting down...")
		cancel()
	}()

	router := mux.NewRouter()
	api.RegisterUserRouters(router)
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "PATCH", "POST", "OPTIONS", "DELETE"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "x-datadog-trace-id", "x-datadog-parent-id", "x-datadog-sampling-priority", "x-datadog-sampled", "x-datadog-origin"}),
	)
	gzip, err := gziphandler.GzipHandlerWithOpts(gziphandler.MinSize(10_000_000))
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Addr:        ":" + strconv.Itoa(port),
		Handler:     cors(gzip(router)),
		ReadTimeout: 2 * time.Minute,
	}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("listening at http://127.0.0.1:%v", port)
	server.ListenAndServe()
}
