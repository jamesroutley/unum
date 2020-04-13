package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jamesroutley/unum/domain"
	"github.com/jamesroutley/unum/server"
	"github.com/jamesroutley/unum/unumpb"
)

var configFile = "config.json"

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
	srv := server.Server{Config: config}
	twirpHandler := unumpb.NewUnumServer(srv, nil)
	log.Println("starting server")
	http.ListenAndServe(":8080", twirpHandler)
}

func loadConfig() (*domain.Config, error) {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	config := &domain.Config{}
	if err := json.Unmarshal(bytes, config); err != nil {
		return nil, err
	}
	return config, nil
}

func WithAuth(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO
		base.ServeHTTP(w, r)
	})
}
