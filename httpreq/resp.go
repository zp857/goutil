package httpreq

import (
	"github.com/imroc/req/v3"
	"github.com/zp857/goutil/constants"
)

func GetHeaderString(resp *req.Response) (headerString string) {
	headerMap := map[string]string{}
	for k := range resp.Header {
		if k != constants.SetCookieHeader {
			headerMap[k] = resp.Header.Get(k)
		}
	}
	for _, ck := range resp.Cookies() {
		headerMap[constants.SetCookieHeader] += ck.String() + ";"
	}
	for k, v := range headerMap {
		headerString += k + ": " + v + "\n"
	}
	return headerString
}

func GetHeaderMap(resp *req.Response) (headerMap map[string][]string) {
	headerMap = map[string][]string{}
	for k := range resp.Header {
		if k != constants.SetCookieHeader {
			headerMap[k] = append(headerMap[k], resp.Header.Get(k))
		}
	}
	for _, ck := range resp.Cookies() {
		headerMap[constants.SetCookieHeader] = append(headerMap[constants.SetCookieHeader], ck.String())
	}
	return headerMap
}
