package main

import (
	"fmt"
	. "github.com/Kumengda/easyChromedp/runtime"
	"github.com/Kumengda/easyChromedp/template"
	"github.com/chromedp/chromedp"
)

func main() {

	templates, err := template.NewChromedpTemplates(
		"http://127.0.0.1:8089/WebGoat/start.mvc",
		10,
		true,
		1,
		map[string]interface{}{"Cookie": "JSESSIONID=rYbGaUnEzJk1n_vJvC8qVpFg7EeQc0lrQtMbjBSa"},
		chromedp.Flag("headless", false),
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
		MainInsp.Print(Json(v))
	}
}
