package headless

import (
	"fmt"
)

func GetHeaderString(headers map[string]interface{}) (headerString string) {
	for k, v := range headers {
		headerString += fmt.Sprintf("%v: %v\n", k, v)
	}
	return headerString
}

func GetHeaderMap(headers map[string]interface{}) (headerMap map[string][]string) {
	for k, v := range headers {
		headerMap[k] = append(headerMap[k], fmt.Sprintf("%v", v))
	}
	return headerMap
}
