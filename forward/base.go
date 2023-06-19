package forward

import (
	"bytes"
	"net/http"
)

// webhooks转发

// CreatePostRequest 创建POST请求
func CreatePostRequest(url string, body map[string]interface{}) (*http.Request, error) {
	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
