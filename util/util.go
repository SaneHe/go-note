package util

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/**
 * @Description: http 请求 client
 */
var client = http.Client{Timeout: 10 * time.Second}

/**
 * @Description: get 方式 http 请求
 * @param url
 * @return []byte
 */
func HttpGet( url string) []byte {
	return DoRequest("GET", url, nil, nil)
}

/**
 * @Description: post 方式 http 请求
 * @param data
 * @param headers
 * @param url
 * @return []byte
 */
func HttpPost(data interface{}, headers map[string]string, url string) []byte {
	return DoRequest("POST", url, data, headers)
}

/**
 * @Description: 执行请求
 * @param method
 * @param url
 * @param data
 * @param headers
 * @return []byte
 */
func DoRequest(method string, url string, data interface{}, headers map[string]string) []byte {

	jsonData, error := json.Marshal(data)
	if error != nil {
		panic(error)
	}

	req, error := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if error != nil {
		panic(error)
	}
	defer req.Body.Close()

	if headers != nil && len(headers) > 0 {
		for index, val := range headers {
			req.Header.Add(index, val)
		}
	}

	logrus.Debug(string(jsonData))
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	//result, _ := httputil.DumpResponse(resp, true)
	return result
}

/**
 * @Description: 根据参数生成标准的url
 * @param urls
 * @return string
 */
func BuildQuery(urls ...string) string {
	params := make([]string, 0, len(urls))

	for _, v := range urls {
		params = append(params, strings.Trim(v, "/"))
	}

	return strings.ToLower(strings.Join(params, "/"))
}