package model

import (
	"io"
	"log"
	"net/http"
	nurl "net/url"
	"time"
)

func TestConnect(url string) (time.Duration, error) {
	now := time.Now()
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	u, err := nurl.Parse(url)
	if err != nil {
		return 0, err
	}
	r, err := client.Do(&http.Request{URL: u})
	if err != nil {
		log.Printf("get %s error: %s", url, err)
		return 0, err
	}
	defer r.Body.Close()
	io.Copy(io.Discard, r.Body)
	return time.Since(now), nil
}
