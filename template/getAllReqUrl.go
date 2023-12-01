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
)

func getWebsiteAllReq(timeout int, websites string, printLog bool, option ...chromedp.ExecAllocatorOption) ([]string, error) {
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
		chromedp.Navigate(websites),
	)

	if err != nil {
		myChrome.Close()
		return nil, err
	}
	myChrome.Close()

	return allReqUrl, nil
}

func getWebsiteAllReqWithsameOrigin(timeout int, websites string, printLog bool, option ...chromedp.ExecAllocatorOption) ([]string, error) {
	allReqUrl, err := getWebsiteAllReq(timeout, websites, printLog, option...)
	if err != nil {
		return nil, err
	}
	sameOriginUrl := sameOriginUrlFilter(websites, allReqUrl)
	return utils.RemoveDuplicateStrings(sameOriginUrl), nil
}

func getWebsiteAllHrefByJs(timeout int, websites string, printLog bool, option ...chromedp.ExecAllocatorOption) ([]string, error) {
	_, err := url.Parse(websites)
	if err != nil {
		return nil, err
	}
	var allOnclickUrl []string
	parse, err := url.Parse(websites)
	if err != nil {
		return nil, err
	}
	host := parse.Host
	scheme := parse.Scheme
	var onclickUrl []string
	myChrome, err := chrome.NewChromeWithTimout(
		timeout,
		option...,
	)
	if err != nil {
		myChrome.Close()
		return nil, err
	}
	err = myChrome.RunWithOutListen(
		chromedp.Navigate(websites),
		chromedp.Evaluate(jsCode.GetAllOnclickUrl, &onclickUrl),
	)

	if err != nil {
		myChrome.Close()
		return nil, err
	}
	for _, u := range onclickUrl {
		//if strings.Contains(u, "?") {
		//	u = u[:strings.Index(u, "?")]
		//}
		if strings.HasPrefix(u, "//") {
			allOnclickUrl = append(allOnclickUrl, scheme+":"+u)
			continue
		}
		if strings.HasPrefix(u, "/") {
			allOnclickUrl = append(allOnclickUrl, scheme+"://"+host+u)
			continue
		}
		if strings.HasPrefix(u, "http") {
			allOnclickUrl = append(allOnclickUrl, u)
			continue
		}
	}
	myChrome.Close()
	return utils.RemoveDuplicateStrings(allOnclickUrl), nil
}

func getWebsiteAllHrefByJsWithSameOrigin(timeout int, websites string, printLog bool, option ...chromedp.ExecAllocatorOption) ([]string, error) {
	allHref, err := getWebsiteAllHrefByJs(timeout, websites, printLog, option...)
	if err != nil {
		return nil, err
	}
	sameOriginUrl := sameOriginUrlFilter(websites, allHref)
	return sameOriginUrl, nil
}
