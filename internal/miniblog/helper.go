package miniblog

import (
	"os"
	"syscall"
	"time"

	"github.com/sjxiang/miniblog/internal/miniblog/store"
	"github.com/sjxiang/miniblog/internal/pkg/db"
	"github.com/sjxiang/miniblog/internal/pkg/log"
)

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             env.LogLevel,
		Format:            env.LogFormat,
		OutputPaths:       env.LogOutputPaths,
	}
}

// 定义要监听的目标信号
var signals = []os.Signal{
	syscall.SIGKILL, // kill -9 pid，该信号不能被捕获 
	syscall.SIGINT,  // kill -2，等同于 CTRL + c
	syscall.SIGTERM, // kill pid
	syscall.SIGQUIT, // ctrl + \
}


// initStore 读取 db 配置，创建 gorm.DB 实例，并初始化 miniblog store 层
func initStore() error {	
	dbOptions := &db.MySQLOptions{
		Host: env.DBHost,
		UserName: env.DBUsername,
		Password: env.DBPassword,
		Database: env.DBDatabase,
		MaxIdleConnections: env.DBMaxIdleConnections,
		MaxOpenConnections: env.DBMaxOpenConnections,
		MaxConnectionLifeTime: time.Duration(env.DBMaxConnectionLifeTime),
		LogLevel: env.DBLogLevel,
	}

	ins, err := db.NewMySQL(dbOptions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)

	return nil 
}