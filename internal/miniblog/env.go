package miniblog

import (
	"github.com/spf13/viper"

	"github.com/sjxiang/miniblog/internal/pkg/log"
)

type Env struct {
	AppEnv         string   `mapstructure:"APP_ENV"`

	RunMode        string   `mapstructure:"RUN_MODE"`
	Addr           string   `mapstructure:"ADDR"`

	DBHost                  string
	DBUsername              string
	DBPassword              string
	DBDatabase              string
	DBMaxIdleConnections    int
	DBMaxOpenConnections    int 
	DBMaxConnectionLifeTime int
	DBLogLevel              int
	
	LogLevel       string   `mapstructure:"LOG_LEVEL"`
	LogFormat      string   `mapstructure:"LOG_FORMAT"`
	LogOutputPaths []string `mapstructure:"LOG_OUTPUT_PATHS"`


}

func NewEnv(filePath string) *Env {
	env := Env{}

	// 从命令行选项指定的配置文件中读取
	viper.SetConfigFile(filePath) //  ".env" && "./.env" 都行

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalw("Can't find the file .env: ", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalw("Environment Can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Infow("The App is running in development env")
	}

	return &env
}
