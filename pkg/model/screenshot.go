package model

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

func TakeScreenshot(url string, wait time.Duration) ([]byte, error) {
	var buf []byte

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithErrorf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	var err error
	width, height := 640, 480
	err = chromedp.Run(ctx, chromedp.Tasks{
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(wait),
		chromedp.CaptureScreenshot(&buf),
	})

	return buf, err
}
