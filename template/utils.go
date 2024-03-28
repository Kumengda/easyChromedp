package template

import (
	"net/url"
	"strings"
)

var typeList = []string{
	"button",
	"checkbox",
	"color",
	"radio",
	"range",
	"reset",
	"search",
	"submit",
}

func sameOriginUrlFilter(websites string, allUrl interface{}) []string {
	var sameOriginPrefix []string
	var sameOriginUrl []string
	parse, _ := url.Parse(websites)
	sameOriginPrefix = append(sameOriginPrefix, parse.Scheme+"://"+parse.Host)
	switch allUrl.(type) {
	case []JsRes:
		for _, v := range allUrl.([]JsRes) {
			for _, v1 := range sameOriginPrefix {
				if strings.HasPrefix(v.Url, v1) {
					sameOriginUrl = append(sameOriginUrl, v.Url)
				}
			}
		}
	case string:
		for _, v := range allUrl.([]string) {
			for _, v1 := range sameOriginPrefix {
				if strings.HasPrefix(v, v1) {
					sameOriginUrl = append(sameOriginUrl, v)
				}
			}
		}
	}

	return sameOriginUrl
}
