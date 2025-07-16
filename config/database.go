package config

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDBConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "hope",
		Password: "123456", // 实际使用时应该从环境变量获取
		DBName:   "hope_blog",
	}
}
