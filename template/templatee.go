package template

import (
	"errors"
	"github.com/chromedp/chromedp"
)

type ChromedpTemplates struct {
	websites string
	timeout  int
	printLog bool
	options  []chromedp.ExecAllocatorOption
}

func NewChromedpTemplates(target string, timeout int, printLog bool, option ...chromedp.ExecAllocatorOption) (*ChromedpTemplates, error) {
	if target == "" || timeout == 0 {
		return nil, errors.New("target and timeout must provide")
	}
	return &ChromedpTemplates{
		websites: target,
		timeout:  timeout,
		printLog: printLog,
		options:  option,
	}, nil
}

func (c *ChromedpTemplates) GetWebsiteAllReq() ([]string, error) {
	allResUrls, err := getWebsiteAllReq(c.timeout, c.websites, c.printLog, c.options...)
	if err != nil {
		return nil, err
	}
	return allResUrls, nil
}

func (c *ChromedpTemplates) GetWebsiteAllReqWithSameOrigin() ([]string, error) {
	allResultUrlsWithSameOrigin, err := getWebsiteAllReqWithsameOrigin(c.timeout, c.websites, c.printLog, c.options...)
	if err != nil {
		return nil, err
	}
	return allResultUrlsWithSameOrigin, err
}

func (c *ChromedpTemplates) GetWebsiteAllHrefByJs() ([]string, error) {
	allOnclickHref, err := getWebsiteAllHrefByJs(c.timeout, c.websites, c.printLog, c.options...)
	if err != nil {
		return nil, err
	}
	return allOnclickHref, err
}
func (c *ChromedpTemplates) GetWebsiteAllHrefByJsWithSameOrigin() ([]string, error) {
	allOnclickHref, err := getWebsiteAllHrefByJsWithSameOrigin(c.timeout, c.websites, c.printLog, c.options...)
	if err != nil {
		return nil, err
	}
	return allOnclickHref, err
}
