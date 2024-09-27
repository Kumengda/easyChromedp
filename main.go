package main

import (
	"fmt"
	"github.com/Kumengda/easyChromedp/chrome"
	"github.com/Kumengda/easyChromedp/template"
	"github.com/chromedp/chromedp"
	"time"
)

func main() {
	myChrome, err := chrome.NewChrome(
		chromedp.Flag("headless", false),
		//chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
	)

	templates, err := template.NewChromedpTemplates(
		true,
		10*time.Second,
		2*time.Second,
		map[string]interface{}{"Cookie": "JSESSIONID=iU6w1WqJJSYb_fQE6yFDYt5jbgb5vtMK0PjT6-fT"},
		myChrome,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	origin, err := templates.GetWebsiteAllHrefByJs("http://127.0.0.1:8765")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range origin {
		fmt.Println(v)
	}
	time.Sleep(5 * time.Second)
	myChrome.Close()

}
