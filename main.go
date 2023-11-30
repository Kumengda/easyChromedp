package main

import (
	"fmt"
	"github.com/Kumengda/easyChromedp/template"
	"github.com/chromedp/chromedp"
)

func main() {
	templates, err := template.NewChromedpTemplates(
		"https://help.mail.163.com/",
		10,
		false,
		chromedp.Flag("headless", true),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	origin, err := templates.GetWebsiteAllHrefByJsWithSameOrigin()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range origin {
		fmt.Println(v)
	}

	origin2, err := templates.GetWebsiteAllReqWithSameOrigin()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range origin2 {
		fmt.Println(v)
	}
}
