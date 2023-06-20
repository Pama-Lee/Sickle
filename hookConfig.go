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

func Test() {
	files, err := getConfigFiles("config")
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

		// fmt.Println(config)

		// 检查配置文件中的 url 是否有效
		if config.URL == "" {
			log.Info("Webhook URL is empty, skip")
			newUUID := tool.GenerateUUID()
			url := Config.Server.Host + ":" + strconv.Itoa(Config.Server.Port) + "/webhooks/" + newUUID
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
