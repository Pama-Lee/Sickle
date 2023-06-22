package server

import (
	"Sickle/log"
	"Sickle/store"
	"Sickle/tool"
	"bytes"
	"encoding/json"
	"io"
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
	j := make(map[string]interface{})
	err := c.BindJSON(&j)

	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// 将headers写入j
	headers := c.Request.Header
	for k, v := range headers {
		j["sickle.headers."+k] = v[0]
	}

	// 生成json
	jsonEn := genJson(j)

	// 获取 config
	configEn := store.UUID2Config[uuid]
	// 匹配 event
	events := configEn.Source.Config.Events
	for _, event := range events {
		if event.Name == jsonEn[event.Key] {
			for _, destination := range event.Destinations {
				// 获取destination
				for _, destinationEn := range configEn.Destinations {
					if destinationEn.Name == destination {
						// 转发
						// 获取url
						url := destinationEn.Config.WebhookURL
						// 获取data
						data := destinationEn.Config.Data
						newData := BindData(data, jsonEn)
						log.Debug("data: ", newData)
						// 转发
						sendForward(url, newData)

					}
				}
			}
			break
		}
	}

	// 获取name
	name := configEn.Name

	// 输出到/config/uuid_json.json
	errS := tool.SaveConfigFile("./request/"+name+"_"+uuid+".json", jsonEn)
	if errS != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})

}

// genJson
func genJson(json map[string]interface{}) map[string]interface{} {
	// 结果
	result := make(map[string]interface{})

	// 遍历json
	for k, v := range json {
		if vMap, ok := v.(map[string]interface{}); ok {
			for k2, v2 := range genJson(vMap) {
				result[k+"."+k2] = v2
			}
		} else {
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(resp.Body)

	// 处理响应
	log.Debug("resp: ", resp)

}

// ReadCloser NewReadCloser
type ReadCloser struct {
	*string
}

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
func BindData(data map[string]interface{}, json map[string]interface{}) map[string]interface{} {
	newData := make(map[string]interface{})

	for k, v := range data {
		newData[k] = v
	}

	// 遍历 data
	for k, v := range data {
		// 如果是 map[string]interface{}
		if vMap, ok := v.(map[string]interface{}); ok {
			// 递归调用，并将结果存储在新变量中
			newVMap := BindData(vMap, json)
			newData[k] = newVMap
		} else {
			// 获取 ${} 中的值
			if str, ok := v.(string); ok {
				keyValue, err := tool.GetKey(str)
				if err != nil {
					log.Error(err)
					continue
				}

				if keyValue != "" {
					// 替换
					newData[k] = json[keyValue]
				}
			}
		}
	}

	return newData
}
