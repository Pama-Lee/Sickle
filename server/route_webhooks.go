package server

import (
	"Sickle/channel"
	"Sickle/log"
	"Sickle/store"
	"Sickle/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
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
						dataList := destinationEn.Config.Data

						// 匹配event
						var data map[string]interface{}
						for _, dataEn := range dataList {
							if dataEn.Event == jsonEn[event.Key] {
								data = dataEn.Data
								break
							}
						}
						// 如果data为空
						if data == nil {
							log.Warn(config.Name, " ", destinationEn.Name, event.Name, " data is nil")
							continue
						}
						newData := BindData(data, jsonEn)
						log.Debug("data: ", newData)
						// 获取是否有headers
						headers := destinationEn.Config.Headers

						// 检查Method是否存在
						if destinationEn.Config.Method == "" {
							destinationEn.Config.Method = "POST"
						}

						// 构建DestinationQueueConfig
						destinationQueueConfig := channel.DestinationQueueConfig{
							WebhookURL: url,
							Method:     destinationEn.Config.Method,
							Data:       newData,
							Headers:    headers,
						}
						// 添加到队列
						channel.AddDestinationQueue(destinationQueueConfig)
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

	for k, v := range data {
		if vMap, ok := v.(map[string]interface{}); ok {
			newVMap := BindData(vMap, json)
			newData[k] = newVMap
		} else {
			if str, ok := v.(string); ok {
				keyValue := tool.GetKey(str)
				if keyValue != nil {
					for _, key := range keyValue {
						if replacement, exists := json[key]; exists {
							str = strings.ReplaceAll(str, "${"+key+"}", fmt.Sprintf("%v", replacement))
						}
					}
					newData[k] = str
				}
			}
		}
	}

	return newData
}
