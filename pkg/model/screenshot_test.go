package model_test

import (
	"google-mirror/pkg/model"
	"log"
	"os"
	"testing"
)

func TestTakeScreenshot(t *testing.T) {
	buf, err := model.TakeScreenshot("https://baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile("baidu.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}
