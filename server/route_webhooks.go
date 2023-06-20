package server

import (
	"Sickle/log"
	"Sickle/store"
	"Sickle/tool"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WebhooksRoute 注册路由
func WebhooksRoute(c *gin.RouterGroup) {
	c.Group("/webhooks")
	{
		c.Any("/:uuid", post)
	}
}

// post 接收webhooks请求
func post(c *gin.Context) {
	uuid := c.Param("uuid")
	// 检查uuid是否存在
	config := store.UUID2Config[uuid]

	if config.UUID == "" {
		log.Error("UUID not found")
		c.JSON(404, gin.H{
			"message": "UUID not found",
		})
		return
	}

	// json
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// 生成json
	jsonEn := genJson(json)

	// 获取 config
	configEn := store.UUID2Config[uuid]
	// 匹配 event
	events := configEn.Source.Config.Events
	for _, event := range events {
		// 匹配 event
		if event.Name == jsonEn[event.Key] {
			// 匹配到
			// 获取要转发到的url
			for _, destination := range event.Destinations {
				// 获取destination
				for _, destinationEn := range configEn.Destinations {
					if destinationEn.Name == destination {
						// 转发
						// 获取url
						url := destinationEn.Config.WebhookURL
						// 获取data
						data := destinationEn.Config.Data
						BindData(data, jsonEn)
						log.Info("data: ", data)
						// 转发
						sendForward(url, data)
					}
				}
			}

			break
		}
	}

	// 输出到/config/uuid_json.json
	tool.SaveConfigFile("./config/"+uuid+"_json.json", jsonEn)

	c.JSON(200, gin.H{
		"message": "ok",
	})

}

// get 接收webhooks请求
func get(c *gin.Context) {

}

// genJson
func genJson(json map[string]interface{}) map[string]interface{} {
	// 结果
	result := make(map[string]interface{})

	// 遍历json
	for k, v := range json {
		// 如果是map[string]interface{}
		if vMap, ok := v.(map[string]interface{}); ok {
			// 递归
			for k2, v2 := range genJson(vMap) {
				// 拼接
				result[k+"."+k2] = v2
			}
		} else {
			// 不是map[string]interface{}
			result[k] = v
		}
	}

	return result
}

// sendForward 发送转发请求
func sendForward(url string, data map[string]interface{}) {
	// 发送 http 请求
	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		log.Error(err)
		return
	}

	// 设置header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("From", "Sickle")

	// 设置body
	// 将 data 转为json
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		return
	}
	// 设置负载
	req.Body = ioutil.NopCloser(bytes.NewReader(jsonData))
	req.ContentLength = int64(len(jsonData))

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	log.Info("resp: ", resp)

}

// NewReadCloser
type ReadCloser struct {
	*string
}

// NewReadCloser
func NewReadCloser(s string) *ReadCloser {
	return &ReadCloser{&s}
}

// BindData 绑定数据
/** 例如传入
"data": {
                    "custom": "${data.event.data.value1}",
                    "custom2": "${data.event.data.value2}",
					"custom3": {
						"custom4": "${data.event.data.value3}"
					}
                }
*/
// 将其中的 ${} 替换为 data 中的值
func BindData(data map[string]interface{}, json map[string]interface{}) {

	// 遍历data
	for k, v := range data {
		// 如果是map[string]interface{}
		if vMap, ok := v.(map[string]interface{}); ok {
			// 递归
			BindData(vMap, json)
		} else {
			// 获取${}中的值
			keyValue, err := tool.GetKey(v.(string))
			if err != nil {
				log.Error(err)
				continue
			}
			// 替换
			data[k] = json[keyValue]
		}
	}
}
