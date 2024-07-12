package main

import (
	"flag"
	"google-mirror/pkg/handler"
	"google-mirror/pkg/model"
	"log"
	"net/http"
	"time"
)

var config *string
var addr *string

func init() {
	config = flag.String("config", "config.yml", "config file")
	addr = flag.String("addr", ":8080", "listen address")
	flag.Parse()
}

func main() {

	var cfg *model.Config
	var err error
	if cfg, err = model.LoadConfig(*config); err != nil {
		log.Fatal(err)
	}

	http.Handle("/list", &handler.ListHandler{})
	http.Handle("/api/get", &handler.GetHandler{})
	http.Handle("/api/screenshot", &handler.ScreenshotHandler{
		WaitTime: time.Duration(cfg.ScreenshotTimeout) * time.Second,
	})
	log.Println("Listening on " + *addr)
	http.ListenAndServe(*addr, nil)
}
