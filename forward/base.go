package forward

import (
	"Sickle/log"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// webhooks转发

// CreateJsonPostRequest 创建POST请求
func CreateJsonPostRequest(url string, body map[string]interface{}, headers map[string]string) (*http.Request, error) {
	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("From", "Sickle")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 设置body
	// 将 body 转为json
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// 设置负载
	req.Body = io.NopCloser(bytes.NewReader(jsonData))
	req.ContentLength = int64(len(jsonData))

	return req, nil
}

// CreateJsonGetRequest 创建GET请求
func CreateGetRequest(url string, body map[string]interface{}, headers map[string]string) (*http.Request, error) {
	// 将参数转为url参数
	newUrl := url + "?"
	for k, v := range body {
		newUrl += k + "=" + v.(string) + "&"
	}

	// 创建请求
	req, err := http.NewRequest("GET", newUrl, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("From", "Sickle")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil

}

// SendForward 发送转发请求
func SendForward(method string, url string, body map[string]interface{}, headers map[string]string) (map[string]interface{}, error) {
	// 创建请求
	var req *http.Request
	var err error
	if method == "POST" {
		req, err = CreateJsonPostRequest(url, body, headers)
	} else if method == "GET" {
		req, err = CreateGetRequest(url, body, headers)
	} else {
		log.Error("Method not supported")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 读取返回
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(resp.Body)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析返回
	var respJson map[string]interface{}
	err = json.Unmarshal(respBody, &respJson)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
