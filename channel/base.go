package channel

// From 1.0.0-alpha

import (
	"Sickle/forward"
	"Sickle/log"
)

type DestinationQueueConfig struct {
	WebhookURL string                 `json:"webhook_url"`
	Method     string                 `json:"method"`
	Data       map[string]interface{} `json:"data"`
	// From 1.0.0-alpha
	Headers map[string]string `json:"headers"`
}

// DestinationQueue 维护一个全局队列
var DestinationQueue = make(chan DestinationQueueConfig, 100)

// AddDestinationQueue 添加到队列
func AddDestinationQueue(config DestinationQueueConfig) {
	DestinationQueue <- config
}

// Start 启动一个goroutine
func Start() {
	log.Info("The destination queue is running...")
	go func() {
		for {
			select {
			case config := <-DestinationQueue:
				// 转发
				// 获取url
				url := config.WebhookURL
				// 获取data
				data := config.Data
				// 获取是否有headers
				headers := config.Headers
				// 转发
				sendForward, err := forward.SendForward(config.Method, url, data, headers)
				if err != nil {
					log.Error(err)
				}
				log.Debug("response from ["+url+"]: ", sendForward)
			}
		}
	}()
}

// GetDestinationQueueLen 获取队列长度
func GetDestinationQueueLen() int {
	return len(DestinationQueue)
}
