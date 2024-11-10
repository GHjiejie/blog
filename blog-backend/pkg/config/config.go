package config

import (
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type Config struct {
	LogLevel       string          // 日志级别
	DEV            bool            // 是否为开发环境
	GRPCEndpoint   string          // GRPC服务地址
	HTTPEndpoint   string          // HTTP服务地址
	EngineEndpoint string          // engine组件 endpoint
	EnginePaths    []string        //
	DBConfig       *DBConfig       // 数据库配置
	UserAuth       *UserAuthConfig // 用户 auth 配置
}

type DBConfig struct {
	Type    string   // 数据库类型
	SQLPara *SQLPara // 数据库配置参数
}

type SQLPara struct {
	DriverName string // 数据库驱动
	User       string // 用户名
	NetMode    string // 网络模式
	Addr       string // 数据库地址
	DBName     string // 数据库名
	Para       string // 其他参数
}

type UserAuthConfig struct {
	TokenExpireDuration string // token 过期时间
	SecretKey           string // token 加密密钥
}

func LoadConfig(configFile string) (*Config, error) {
	// https://github.com/go-critic/go-critic/issues/1019.
	// 读取配置文件
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	//这个结构体用于存储配置文件中的配置信息，最后会返回给调用者
	var c Config
	// 可视化转换: https://xuri.me/toml-to-go/.
	err = toml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	// 从环境变量获取敏感配置
	// MySQL Config.
	// if v, ok := os.LookupEnv("MYSQL_ADDRESS"); ok {
	// 	c.DBConfig.SQLPara.Addr = v
	// }

	// if v, ok := os.LookupEnv("MYSQL_DATABASE"); ok {
	// 	c.DBConfig.SQLPara.DBName = v
	// }

	// username, okU := os.LookupEnv("MYSQL_USERNAME")
	// password, okP := os.LookupEnv("MYSQL_PASSWORD")
	// if okU && okP {
	// 	c.DBConfig.SQLPara.User = fmt.Sprintf("%s:%s", username, password)
	// }

	return &c, nil
}
