package httputil

import (
	"github.com/imroc/req/v3"
	"github.com/zp857/goutil/validator"
	"net"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	reg1 = regexp.MustCompile(`(?i)<meta.*?http-equiv=.*?refresh.*?url=(.*?)/?>`)

	reg2 = regexp.MustCompile(`(?i)[window\.]?location[\.href]?.*?=.*?["'](.*?)["']`)

	reg3 = regexp.MustCompile(`(?i)[window\.]?location\.replace\(['"](.*?)['"]\)`)
)

var (
	regHost = regexp.MustCompile(`(?i)https?://(.*?)/`)
)

func JsJump(resp *req.Response) (jumpUrl string) {
	res := regexJsJump(resp)
	if res != "" {
		res = strings.TrimSpace(res)
		res = strings.ReplaceAll(res, "\"", "")
		res = strings.ReplaceAll(res, "'", "")
		res = strings.ReplaceAll(res, "./", "/")
		if strings.HasPrefix(res, "http") {
			matches := regHost.FindAllStringSubmatch(res, -1)
			if len(matches) > 0 {
				var ip net.IP
				if strings.Contains(matches[0][1], ":") {
					ip = net.ParseIP(strings.Split(matches[0][1], ":")[0])
				} else {
					ip = net.ParseIP(matches[0][1])
				}
				if validator.IsInnerIp(ip.String()) {
					baseURL := resp.Request.URL.Host
					res = strings.ReplaceAll(res, matches[0][1], baseURL)
				}
			}
			jumpUrl = res
		} else if strings.HasPrefix(res, "/") {
			// 前缀存在 / 时拼接绝对目录
			baseURL := resp.Request.URL.Scheme + "://" + resp.Request.URL.Host
			jumpUrl = baseURL + res
		} else {
			// 前缀不存在 / 时拼接相对目录
			baseURL := resp.Request.URL.Scheme + "://" + resp.Request.URL.Host + "/" + filepath.Dir(resp.Request.URL.Path) + "/"
			baseURL = strings.ReplaceAll(baseURL, "./", "")
			jumpUrl = baseURL + res
		}
	}
	return
}

func regexJsJump(resp *req.Response) string {
	matches := reg1.FindAllStringSubmatch(resp.String(), -1)
	if len(matches) > 0 {
		// 去除注释的情况
		if !strings.Contains(resp.String(), "<!--\r\n"+matches[0][0]) && !strings.Contains(matches[0][1], "nojavascript.html") && !strings.Contains(resp.String(), "<!--[if lt IE 7]>\n"+matches[0][0]) {
			return matches[0][1]
		}
	}
	body := resp.String()
	if len(body) > 700 {
		body = body[:700]
	}
	matches = reg2.FindAllStringSubmatch(body, -1)
	if len(matches) > 0 {
		return matches[0][1]
	}
	matches = reg3.FindAllStringSubmatch(body, -1)
	if len(matches) > 0 {
		return matches[0][1]
	}
	return ""
}
