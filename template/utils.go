package template

import (
	"net/url"
	"strings"
)

func sameOriginUrlFilter(websites string, allUrl []string) []string {
	var sameOriginPrefix []string
	var sameOriginUrl []string
	parse, _ := url.Parse(websites)
	sameOriginPrefix = append(sameOriginPrefix, parse.Scheme+"://"+parse.Host)
	for _, v := range allUrl {
		for _, v1 := range sameOriginPrefix {
			if strings.HasPrefix(v, v1) {
				sameOriginUrl = append(sameOriginUrl, v)
			}
		}
	}
	return sameOriginUrl
}
