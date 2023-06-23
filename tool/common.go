package tool

import (
	"Sickle/log"
	"encoding/json"
	"os"
	"regexp"

	"github.com/google/uuid"
)

// FileExists 检查文件是否存在
func FileExists(path string) bool {
	// 打开文件
	_, err := os.Open(path)
	// 文件不存在
	if err != nil && os.IsNotExist(err) {
		return false
	}
	// 文件存在
	return true
}

// GenerateUUID 生成UUID
func GenerateUUID() string {
	// 使用谷歌的UUID生成算法
	uuid, err := uuid.New().MarshalText()
	if err != nil {
		return ""
	}
	return string(uuid)
}

// SaveConfigFile 保存配置文件, json
func SaveConfigFile(path string, config interface{}) error {
	// 保存
	file, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		log.Error(err)
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Error(err)
		}
	}(f)

	_, err = f.Write(file)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func GetKey(input string) []string {
	variables := make([]string, 0)

	re := regexp.MustCompile(`\$\{([^}]+)\}`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		variables = append(variables, match[1])
	}

	return variables
}

// IsString 判断是否是字符串
func IsString(str interface{}) bool {
	_, ok := str.(string)
	return ok
}
