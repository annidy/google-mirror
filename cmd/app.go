package main

import (
	"flag"
	"google-mirror/pkg/handler"
	"google-mirror/pkg/model"
	"log"
	"net/http"
	"sync"
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

	listHandler := handler.ListHandler{}

	wg := sync.WaitGroup{}
	for i := 0; i < len(cfg.Mirrors); i++ {
		wg.Add(1)
		go func(i int) {
			m := make(map[string]any)
			defer wg.Done()
			urls := model.ExtractURLs(cfg.Mirrors[i])
			for j := 0; j < len(urls); j++ {
				m[urls[j]] = struct{}{}
			}
			item := handler.ListItem{
				Name:    cfg.Mirrors[i],
				Mirrors: []model.Mirror{},
			}
			for k := range m {
				item.Mirrors = append(item.Mirrors, model.Mirror{
					URL: k,
				})
			}
			item.Mirrors = item.Mirrors[0:3] // TODO: debug
			listHandler.Data.Items = append(listHandler.Data.Items, item)
		}(i)
	}
	wg.Wait()

	http.Handle("/list", &listHandler)
	http.Handle("/api/get", &handler.GetHandler{})
	log.Println("Listening on " + *addr)
	http.ListenAndServe(*addr, nil)
}
