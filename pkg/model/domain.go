package model

import (
	"log"
	"net/http"
	"time"
)

func (m *Mirror) TakeSnapshot() error {
	now := time.Now()
	if _, err := http.Get(m.URL); err != nil {
		log.Printf("get %s error: %s", m.URL, err)
		return err
	}
	m.ConnectTime = time.Since(now)

	return nil
}
