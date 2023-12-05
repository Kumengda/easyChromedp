package main

import (
	"fmt"
	"github.com/Kumengda/easyChromedp/template"
	"github.com/chromedp/chromedp"
)

func main() {
	templates, err := template.NewChromedpTemplates(
		"https://easm.huaun.com:443",
		10,
		true,
		5,
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

	origin2, err := templates.GetWebsiteAllReq()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range origin2 {
		fmt.Println(v)
	}
}
