package template

import (
	"fmt"
	"github.com/Kumengda/easyChromedp/chrome"
	"github.com/Kumengda/easyChromedp/jsCode"
	. "github.com/Kumengda/easyChromedp/runtime"
	"github.com/Kumengda/easyChromedp/utils"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"net/url"
	"strings"
	"time"
)

func getWebsiteAllReq(timeout int, websites string, printLog bool, waitTime int, headers map[string]interface{}, option ...chromedp.ExecAllocatorOption) ([]string, error) {
	_, err := url.Parse(websites)
	if err != nil {
		return nil, err
	}
	var allReqUrl []string
	myChrome, err := chrome.NewChromeWithTimout(
		timeout,
		option...,
	)
	if err != nil {
		myChrome.Close()
		return nil, err
	}
	err = myChrome.RunWithListen(func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			reqUrl := ev.Request.URL
			//if strings.Contains(reqUrl, "?") {
			//	reqUrl = reqUrl[:strings.Index(reqUrl, "?")]
			//}
			if strings.HasPrefix(reqUrl, "http") {
				allReqUrl = append(allReqUrl, reqUrl)
			}
			if printLog {
				MainInsp.Print(LEVEL_DEBUG, Text(fmt.Sprintf("req url %s", ev.Request.URL)))
			}
		}
	},
		network.Enable(),
		network.SetExtraHTTPHeaders(headers),
		chromedp.Navigate(websites),
		chromedp.Sleep(time.Duration(waitTime)*time.Second),
	)

	if err != nil {
		myChrome.Close()
		return nil, err
	}
	myChrome.Close()

	return allReqUrl, nil
}

func getWebsiteAllReqWithsameOrigin(timeout int, websites string, printLog bool, waitTime int, headers map[string]interface{}, option ...chromedp.ExecAllocatorOption) ([]string, error) {
	allReqUrl, err := getWebsiteAllReq(timeout, websites, printLog, waitTime, headers, option...)
	if err != nil {
		return nil, err
	}
	sameOriginUrl := sameOriginUrlFilter(websites, allReqUrl)
	return utils.RemoveDuplicateStrings(sameOriginUrl), nil
}

func getWebsiteAllHrefByJs(timeout int, websites string, printLog bool, headers map[string]interface{}, waitTime int, option ...chromedp.ExecAllocatorOption) ([]JsRes, error) {
	var allOnclickUrl []JsRes
	parse, err := url.Parse(websites)
	if err != nil {
		return nil, err
	}
	host := parse.Host
	scheme := parse.Scheme
	var onclickUrl []string

	var fromDatas []FormDatas
	myChrome, err := chrome.NewChromeWithTimout(
		timeout,
		option...,
	)
	if err != nil {
		myChrome.Close()
		return nil, err
	}
	err = myChrome.RunWithOutListen(
		network.Enable(),
		network.SetExtraHTTPHeaders(headers),
		chromedp.Navigate(websites),
		chromedp.Sleep(time.Duration(waitTime)*time.Second),
		chromedp.Evaluate(jsCode.GetAllOnclickUrl, &onclickUrl),
		chromedp.Evaluate(jsCode.ParseFrom, &fromDatas),
	)

	if err != nil {
		myChrome.Close()
		return nil, err
	}
	onclickUrl = cleanOnclickUrl(onclickUrl)
	onclickUrl = utils.RemoveDuplicateStrings(onclickUrl)
	for _, u := range onclickUrl {
		allOnclickUrl = append(allOnclickUrl, JsRes{
			Url:    parseJsData(u, scheme, host, websites),
			Method: "GET",
			Param:  nil,
		})
	}
	for _, v := range fromDatas {
		var fromUrl string
		isFileUpload := false
		for _, vv := range v.FormData {
			if vv.Type == "file" {
				isFileUpload = true
			}
		}
		if v.Action == "#" || v.Action == "/" || v.Action == "" {
			fromUrl = websites
		} else {
			fromUrl = parseJsData(v.Action, scheme, host, websites)
		}
		allOnclickUrl = append(allOnclickUrl, JsRes{
			Url:          fromUrl,
			Method:       strings.ToUpper(v.Method),
			IsForm:       true,
			Param:        v.FormData,
			IsFileUpload: isFileUpload,
		})
	}
	myChrome.Close()
	return allOnclickUrl, nil
}

func getWebsiteAllHrefByJsWithSameOrigin(timeout int, websites string, printLog bool, headers map[string]interface{}, waitTime int, option ...chromedp.ExecAllocatorOption) ([]string, error) {
	allHref, err := getWebsiteAllHrefByJs(timeout, websites, printLog, headers, waitTime, option...)
	if err != nil {
		return nil, err
	}
	sameOriginUrl := sameOriginUrlFilter(websites, allHref)
	return sameOriginUrl, nil
}

func parseJsData(u, scheme, host, nowUrl string) string {

	//if strings.Contains(u, "?") {
	//	u = u[:strings.Index(u, "?")]
	//}
	if strings.HasPrefix(u, "//") {
		return scheme + ":" + u
	}
	if strings.HasPrefix(u, "/") {
		return scheme + "://" + host + u
	}
	if strings.HasPrefix(u, "#") {
		return nowUrl + u

	}
	if strings.HasPrefix(u, "http") {
		return u
	}
	nowUrl = nowUrl[:strings.LastIndex(nowUrl, "/")]
	return nowUrl + "/" + u
}
func cleanOnclickUrl(target []string) []string {
	var newUrls = []string{}
	for _, v := range target {
		if replace(v, "", "#", "\n", " ") != "" {
			newUrls = append(newUrls, v)
		}
	}
	return newUrls
}

func replace(rawStr string, replaceStr ...string) string {
	for _, v := range replaceStr {
		rawStr = strings.ReplaceAll(rawStr, v, "")
	}
	return rawStr
}
