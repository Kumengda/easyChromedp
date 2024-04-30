package main

import (
	"fmt"
	"github.com/Kumengda/easyChromedp/chrome"
	"time"

	//. "github.com/Kumengda/easyChromedp/runtime"
	"github.com/Kumengda/easyChromedp/template"
	"github.com/chromedp/chromedp"
)

func main() {
	myChrome, err := chrome.NewChrome(
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
	)
	templates, err := template.NewChromedpTemplates(
		true,
		1,
		0,
		map[string]interface{}{"Cookie": "JSESSIONID=iU6w1WqJJSYb_fQE6yFDYt5jbgb5vtMK0PjT6-fT"},
		myChrome,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	origin, err := templates.GetWebsiteAllHrefByJs("https://www.baidu.com")
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
