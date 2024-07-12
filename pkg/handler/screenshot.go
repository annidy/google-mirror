package handler

import (
	"google-mirror/pkg/model"
	"log"
	"net/http"
	"time"
)

type ScreenshotHandler struct {
	WaitTime time.Duration
}

func (h *ScreenshotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	url := r.URL.Query()["url"][0]
	log.Println("handler screenshot", url)

	if b, err := model.TakeScreenshot(url, h.WaitTime); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Max-Age", "86400")
		w.Write(b)
	}
}
