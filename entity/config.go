package entity

// Config 配置文件结构体
type Config struct {
	// 数据库配置
	Database struct {
		// 数据库类型
		Type string `json:"type"`
		// 数据库地址
		Host string `json:"host"`
		// 数据库端口
		Port int `json:"port"`
		// 数据库用户名
		User string `json:"user"`
		// 数据库密码
		Password string `json:"password"`
		// 数据库名称
		Database string `json:"database"`
	} `json:"database"`
	// 服务器配置
	Server struct {
		// 服务器地址
		Host string `json:"host"`
		// 服务器端口
		Port int `json:"port"`
	} `json:"server"`
	// 日志配置
	Log struct {
		// 日志文件路径
		File string `json:"file"`
		// 日志级别
		Level string `json:"level"`
	} `json:"log"`
	// 项目配置
	Project struct {
		// 项目名称
		Name string `json:"name"`
		// 项目版本
		Version string `json:"version"`
		// 项目作者
		Author string `json:"author"`
		// 项目描述
		Description string `json:"description"`
	} `json:"project"`
}
