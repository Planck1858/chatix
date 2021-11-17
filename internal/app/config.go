package app

type Config struct {
	Environment string `yaml:"environment"`
	AppName     string `yaml:"appName"`
	Server      struct {
		Port     string `yaml:"port"`
		BasePath string `yaml:"basePath"`
		TimeOut  int    `yaml:"timeOut"`
	} `yaml:"application"`
	Db struct {
		PostgresConnection string `yaml:"postgresConnection"`
		MaxIdleConnections int    `yaml:"maxIdleConnections"`
		MaxOpenConnections int    `yaml:"maxOpenConnections"`
	} `yaml:"db"`
	Log struct {
		Title     string `yaml:"title"`
		Type      string `yaml:"type"`
		Level     string `yaml:"level"`
		Formatter string `yaml:"formatter"`
	} `yaml:"log"`
}
