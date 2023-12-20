package main

import (
	"fmt"
	"github.com/Kumengda/easyChromedp/template"
	"github.com/chromedp/chromedp"
)

func main() {

	templates, err := template.NewChromedpTemplates(
		"http://127.0.0.1:9999",
		10,
		true,
		5,
		map[string]interface{}{"Cookie": "PHPSESSID=0orsrs37tjva2opouppr9agvn1; security=low;"},
		chromedp.Flag("headless", true),
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	//origin, err := templates.GetWebsiteAllHrefByJsWithSameOrigin()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, v := range origin {
	//	fmt.Println(v)
	//}

	origin2, err := templates.GetWebsiteAllHrefByJs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range origin2 {
		fmt.Println(v)
	}
}
