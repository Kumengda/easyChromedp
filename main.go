package main

import (
	"fmt"
	"github.com/Kumengda/easyChromedp/chrome"
	//. "github.com/Kumengda/easyChromedp/runtime"
	"github.com/Kumengda/easyChromedp/template"
	"github.com/chromedp/chromedp"
)

func main() {
	myChrome, err := chrome.NewChromeWithTimout(
		10,
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
	)

	templates, err := template.NewChromedpTemplates(
		true,
		2,
		map[string]interface{}{"Cookie": "JSESSIONID=iU6w1WqJJSYb_fQE6yFDYt5jbgb5vtMK0PjT6-fT"},
		myChrome,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, can := chromedp.NewContext(myChrome.GetContext())

	origin, err := templates.GetWebsiteAllHrefByJsWithSameOrigin(ctx, "http://127.0.0.1:8765")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range origin {
		fmt.Println(v)
	}
	can()
	myChrome.Close()
}
