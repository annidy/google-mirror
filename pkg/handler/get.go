package handler

import (
	"encoding/json"
	"google-mirror/pkg/model"
	"log"
	"net/http"
	"time"
)

type GetHandler struct {
	URL string
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	decoder.Decode(&params)

	log.Printf("get %s", params)

	url := params["url"]
	var t time.Duration
	var err error
	if t, err = model.TestConnect(url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type R struct {
		Time int    `json:"time"`
		Host string `json:"host"`
	}

	rsp, err := json.Marshal(R{
		Time: int(t.Milliseconds()),
		Host: url,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rsp)
}
