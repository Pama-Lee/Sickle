package main

import (
	"Sickle/entity"
	"Sickle/log"
	"Sickle/store"
	"Sickle/tool"
	"encoding/json"
	"strconv"

	"os"
	"path/filepath"
)

func hookConfig() {
	files, err := getConfigFiles("webhooks")

	// 检查request文件夹是否存在, 不存在则创建
	if _, err := os.Stat("request"); os.IsNotExist(err) {
		err := os.Mkdir("request", os.ModePerm)
		if err != nil {
			log.Error(err)
		}
	}

	if err != nil {
		log.Error("Error:", err)
		return
	}

	i := 0

	for _, file := range files {
		log.Debug("Processing:", file)
		config, err := readConfigFile(file)
		if err != nil {
			log.Error("Error:", err)
			continue
		}

		i = i + 1

		// 检查配置文件中的 url 是否有效
		if config.URL == "" {
			log.Info("Webhook URL is empty, generating a new one...")
			newUUID := tool.GenerateUUID()

			var url string

			// 如果Port为空或者为80，那么就不需要加端口号
			if Config.Server.Port == 80 || Config.Server.Port == 0 {
				url = Config.Server.Host + "/webhooks/" + newUUID
			} else {
				url = Config.Server.Host + ":" + strconv.Itoa(Config.Server.Port) + "/webhooks/" + newUUID
			}
			config.URL = url
			config.UUID = newUUID
			// 保存配置文件
			err = tool.SaveConfigFile(file, config)
			if err != nil {
				log.Error(err)
				continue
			}
		}

		store.GlobalConfig = append(store.GlobalConfig, *config)
		store.UUID2Config[config.UUID] = *config
	}

	log.Debug("Total:", i)
}

func getConfigFiles(folderPath string) ([]string, error) {
	var files []string

	// 检查文件夹是否存在, 不存在则创建
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			return nil, err
		}

		// 创建一个空的示例配置文件
		exampleConfig := entity.Webhook{
			Name: "Example",
			Source: entity.Source{
				Type: "github",
				Config: entity.SourceConfig{
					Events: []entity.Event{
						{
							Name:         "push",
							Key:          "sickle.headers.X-GitHub-Event",
							Destinations: []string{"DingTalk", "WeCom"},
						},
						{
							Name:         "pull_request",
							Key:          "sickle.headers.X-GitHub-Event",
							Destinations: []string{"DingTalk"},
						},
					},
				},
			},
			Destinations: []entity.Destination{
				{
					Name: "DingTalk",
					Type: "dingtalk",
					Config: entity.DestinationConfig{
						WebhookURL: "https://oapi.dingtalk.com/robot/send?access_token=xxx",
						Data: map[string]interface{}{
							"msgtype": "markdown",
							"markdown": map[string]interface{}{
								"title": "Sickle",
								"text": "### Sickle\n" +
									"**{{sickle.headers.X-GitHub-Event}}**\n" +
									"**{{action}}**\n",
							},
						},
					},
				},
				{
					Name: "WeCom",
					Type: "WeCom",
					Config: entity.DestinationConfig{
						WebhookURL: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx",
						Data: map[string]interface{}{
							"msgtype": "text",
							"text": map[string]interface{}{
								"content": "Sickle\n" +
									"{{sickle.headers.X-GitHub-Event}}\n" +
									"{{action}}\n",
							},
						},
					},
				},
			},
		}

		err = tool.SaveConfigFile(filepath.Join(folderPath, "example.json"), &exampleConfig)
		if err != nil {
			return nil, err
		}
		log.Info("Example config file created, please edit it and restart Sickle.")

	}

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func readConfigFile(filePath string) (*entity.Webhook, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config entity.Webhook
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
