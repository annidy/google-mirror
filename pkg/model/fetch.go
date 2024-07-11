package model

import (
	"io"
	"net/http"
	nurl "net/url"
	"regexp"
)

func ExtractURLs(url string) []string {
	u, err := nurl.Parse(url)
	if err != nil {
		return nil
	}
	rsp, err := http.Get(u.String())
	if err != nil {
		return nil
	}
	defer rsp.Body.Close()
	s, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil
	}
	links := ExtractLinks(string(s))
	return links
}

func ExtractLinks(doc string) []string {
	urlRegex := regexp.MustCompile(`(?i)\b((?:https?):\/\/[-A-Z0-9+&@#\/%?=~_|!:,.;]*[-A-Z0-9+&@#\/%=~_|])\b`)
	matches := urlRegex.FindAllString(doc, -1)
	rets := []string{}

	for _, m := range matches {
		u, err := nurl.Parse(m)
		if err != nil {
			continue
		}
		if u.Path != "" && u.Path != "/" {
			continue
		}
		if len(u.Query()) > 0 {
			continue
		}
		if len(u.Fragment) > 0 {
			continue
		}
		rets = append(rets, u.String())
	}
	return rets
}
