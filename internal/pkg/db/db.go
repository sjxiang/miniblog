package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
)

// MySQLOptions 定义 MySQL 数据库的选项
type MySQLOptions struct {
	Host                  string
	UserName              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int 
}


// DSN 从 MySQLOptions 返回 DSN
func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&&multiStatements=%t&loc=%s`,
		o.UserName,
		o.Password,
		o.Host,
		o.Database,
		true,
		true,
		"Local")  
}

// NewMySQL 使用给定的选项创建一个新的 gorm 数据库实例
func NewMySQL(opts *MySQLOptions) (*gorm.DB, error) {
	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}

	db, err := gorm.Open(mysql.Open(opts.DSN()), 
		&gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
			SkipDefaultTransaction: true,  // 关闭默认事务
			PrepareStmt: true,             // 缓存预编译语句
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}


	sqlDB.SetMaxIdleConns(opts.MaxOpenConnections)        // 设置连接池的最大连接数
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)  // 连接的生命周期
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)        // 空闲队列的最大连接数

	return db, nil 
}