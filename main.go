package main

import (
	"fmt"
	"github.com/Kumengda/easyChromedp/chrome"
	"github.com/chromedp/chromedp"
)

func main() {

	//templates, err := template.NewChromedpTemplates(
	//	"https://www.baidu.com",
	//	10,
	//	true,
	//	5,
	//	map[string]interface{}{"Cookie": "PHPSESSID=0orsrs37tjva2opouppr9agvn1; security=low;"},
	//	chromedp.Flag("headless", false),
	//	chromedp.DisableGPU,
	//	chromedp.NoDefaultBrowserCheck,
	//)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	////origin, err := templates.GetWebsiteAllHrefByJsWithSameOrigin()
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}
	////for _, v := range origin {
	////	fmt.Println(v)
	////}
	//
	//origin2, err := templates.GetWebsiteAllHrefByJs()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, v := range origin2 {
	//	MainInsp.Print(Json(v))
	//}

	chro, _ := chrome.NewChromeWithTimout(10,
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
	)
	err := chro.RunWithOutListen(
		chromedp.Navigate("about:blank"),
	)
	if err != nil {
		fmt.Print(err)
		return
	}
	chromedp.NewRemoteAllocator(chro.GetContext(), "")

}
