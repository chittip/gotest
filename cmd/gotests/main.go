package main

import (
	"log"
	"net/http"

	"github.com/chittip/gonews/pkg/model"

	"github.com/chittip/gonews/pkg/app"
)

const (
	port     = ":8080"
	mongoURL = "mongodb://127.0.0.1:27017"
)

func main() {
	mux := http.NewServeMux()
	// if mux {

	// }
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatal("can not init model; %w", err)
	}
	http.ListenAndServe(port, mux)
}
