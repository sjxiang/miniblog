package miniblog

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv string `mapstructure:"APP_ENV"`
}

func NewEnv(filePath string) *Env {
	env := Env{}

	// 从命令行选项指定的配置文件中读取
	viper.SetConfigFile(filePath) //  ".env" && "./.env" 都行

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Environment Can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
