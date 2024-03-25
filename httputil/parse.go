package httputil

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// ParseQueryParams 解析 GET 请求的 URL 查询参数
func ParseQueryParams(urlStr string) (map[string]string, error) {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	for key, values := range parsedUrl.Query() {
		queryParams[key] = values[0] // 取第一个值
	}
	return queryParams, nil
}

// ParsePostData 解析 POST 请求的 PostData
func ParsePostData(postData, contentType string) (map[string]interface{}, error) {
	switch {
	case strings.Contains(contentType, "application/x-www-form-urlencoded"):
		return parseFormPostData(postData)
	case strings.Contains(contentType, "application/json"):
		return parseJSONPostData(postData)
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}
}

func parseFormPostData(data string) (map[string]interface{}, error) {
	parsedData, err := url.ParseQuery(data)
	if err != nil {
		return nil, err
	}
	dataMap := make(map[string]interface{})
	for key, values := range parsedData {
		dataMap[key] = values[0] // 取第一个值
	}
	return dataMap, nil
}

func parseJSONPostData(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
