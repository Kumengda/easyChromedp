package template

import (
	"errors"
	"github.com/chromedp/chromedp"
)

type ChromedpTemplates struct {
	websites string
	timeout  int
	waitTime int
	printLog bool
	headers  map[string]interface{}
	options  []chromedp.ExecAllocatorOption
}

func NewChromedpTemplates(target string, timeout int, printLog bool, waitTime int, headers map[string]interface{}, option ...chromedp.ExecAllocatorOption) (*ChromedpTemplates, error) {
	if target == "" || timeout == 0 {
		return nil, errors.New("target and timeout must provide")
	}
	if waitTime >= timeout {
		return nil, errors.New("waitTime不可大于等于timeout")
	}
	return &ChromedpTemplates{
		waitTime: waitTime,
		websites: target,
		timeout:  timeout,
		printLog: printLog,
		options:  option,
		headers:  headers,
	}, nil
}

func (c *ChromedpTemplates) GetWebsiteAllReq() ([]string, error) {
	allResUrls, err := getWebsiteAllReq(c.timeout, c.websites, c.printLog, c.waitTime, c.headers, c.options...)
	if err != nil {
		return nil, err
	}
	return allResUrls, nil
}

func (c *ChromedpTemplates) GetWebsiteAllReqWithSameOrigin() ([]string, error) {
	allResultUrlsWithSameOrigin, err := getWebsiteAllReqWithsameOrigin(c.timeout, c.websites, c.printLog, c.waitTime, c.headers, c.options...)
	if err != nil {
		return nil, err
	}
	return allResultUrlsWithSameOrigin, err
}

func (c *ChromedpTemplates) GetWebsiteAllHrefByJs() ([]JsRes, error) {
	allOnclickHref, err := getWebsiteAllHrefByJs(c.timeout, c.websites, c.printLog, c.headers, c.waitTime, c.options...)
	if err != nil {
		return nil, err
	}
	return allOnclickHref, err
}
func (c *ChromedpTemplates) GetWebsiteAllHrefByJsWithSameOrigin() ([]string, error) {
	allOnclickHref, err := getWebsiteAllHrefByJsWithSameOrigin(c.timeout, c.websites, c.printLog, c.headers, c.waitTime, c.options...)
	if err != nil {
		return nil, err
	}
	return allOnclickHref, err
}
