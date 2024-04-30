package template

import (
	"context"
	"errors"
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

type ChromedpTemplates struct {
	waitTime int
	printLog bool
	headers  map[string]interface{}
	chrome   *chrome.Chrome
}

func NewChromedpTemplates(printLog bool, waitTime int, headers map[string]interface{}, chrome *chrome.Chrome) (*ChromedpTemplates, error) {
	if chrome.GetTimeout() == 0 {
		return nil, errors.New("timeout must provide")
	}
	if waitTime >= chrome.GetTimeout() {
		return nil, errors.New("waitTime不可大于等于timeout")
	}
	return &ChromedpTemplates{
		waitTime: waitTime,
		printLog: printLog,
		headers:  headers,
		chrome:   chrome,
	}, nil
}

func (t *ChromedpTemplates) GetWebsiteAllReq(ctx context.Context, websites string) ([]string, error) {
	_, err := url.Parse(websites)
	if err != nil {
		return nil, err
	}
	var allReqUrl []string

	err = t.chrome.RunWithListen(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			reqUrl := ev.Request.URL
			//if strings.Contains(reqUrl, "?") {
			//	reqUrl = reqUrl[:strings.Index(reqUrl, "?")]
			//}
			if strings.HasPrefix(reqUrl, "http") {
				allReqUrl = append(allReqUrl, reqUrl)
			}
			if t.printLog {
				MainInsp.Print(LEVEL_DEBUG, Text(fmt.Sprintf("req url %s", ev.Request.URL)))
			}
		}
	},
		network.Enable(),
		network.SetExtraHTTPHeaders(t.headers),
		chromedp.Navigate(websites),
		chromedp.Sleep(time.Duration(t.waitTime)*time.Second),
	)
	return allReqUrl, nil
}

func (t *ChromedpTemplates) GetWebsiteAllReqWithsameOrigin(ctx context.Context, websites string) ([]string, error) {
	allReqUrl, err := t.GetWebsiteAllReq(ctx, websites)
	if err != nil {
		return nil, err
	}
	sameOriginUrl := sameOriginUrlFilter(websites, allReqUrl)
	return utils.RemoveDuplicateStrings(sameOriginUrl), nil
}

func (t *ChromedpTemplates) GetWebsiteAllHrefByJs(ctx context.Context, websites string) ([]JsRes, error) {
	var allOnclickUrl []JsRes
	parse, err := url.Parse(websites)
	if err != nil {
		return nil, err
	}
	host := parse.Host
	scheme := parse.Scheme
	var onclickUrl []string

	var fromDatas []FormDatas
	err = t.chrome.RunWithOutListen(
		ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(t.headers),
		chromedp.Navigate(websites),
		chromedp.Sleep(time.Duration(t.waitTime)*time.Second),
		chromedp.Evaluate(jsCode.GetAllOnclickUrl, &onclickUrl),
		chromedp.Evaluate(jsCode.ParseFrom, &fromDatas),
	)

	if err != nil {
		return nil, err
	}
	onclickUrl = cleanOnclickUrl(onclickUrl)
	onclickUrl = utils.RemoveDuplicateStrings(onclickUrl)
	for _, u := range onclickUrl {
		allOnclickUrl = append(allOnclickUrl, JsRes{
			Url:    parseJsData(u, scheme, host, websites, false),
			Method: "GET",
			Param:  nil,
		})
	}
	for _, v := range fromDatas {
		var fromUrl string
		var newFormData []FormData
		isFileUpload := false
		for _, vv := range v.FormData {
			if vv.Name == "" || !checkInputType(vv.Type) {
				continue
			}
			if vv.Type == "file" {
				isFileUpload = true
			}
			newFormData = append(newFormData, vv)
		}
		if v.Action == "#" || v.Action == "/" || v.Action == "" {
			fromUrl = websites
		} else {
			fromUrl = parseJsData(v.Action, scheme, host, websites, true)
		}
		allOnclickUrl = append(allOnclickUrl, JsRes{
			Url:          fromUrl,
			Method:       strings.ToUpper(v.Method),
			IsForm:       true,
			Param:        newFormData,
			IsFileUpload: isFileUpload,
		})
	}
	return allOnclickUrl, nil
}

func (t *ChromedpTemplates) GetWebsiteAllHrefByJsWithSameOrigin(ctx context.Context, websites string) ([]string, error) {
	allHref, err := t.GetWebsiteAllHrefByJs(ctx, websites)
	if err != nil {
		return nil, err
	}
	sameOriginUrl := sameOriginUrlFilter(websites, allHref)
	return sameOriginUrl, nil
}
