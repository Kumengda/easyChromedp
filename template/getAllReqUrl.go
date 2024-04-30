package template

import (
	"strings"
)

func parseJsData(u, scheme, host, nowUrl string, isForm bool) string {

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
	if isForm && strings.Contains(nowUrl, "#") {
		nowUrl = nowUrl[:strings.Index(nowUrl, "#")]
	}
	if strings.HasSuffix(nowUrl, "/") {
		nowUrl = nowUrl[:strings.LastIndex(nowUrl, "/")]
	}
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

func checkInputType(t string) bool {
	check := strings.ToLower(t)
	for _, v := range typeList {
		if check == v {
			return false
		}
	}
	return true
}
