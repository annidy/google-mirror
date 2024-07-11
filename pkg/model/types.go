package model

import "time"

type Mirror struct {
	URL         string
	ConnectTime time.Duration
	Snapshot    []byte
}
