package database

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"vvvstore/internal/app/model"
)

type Config struct {
	User string
	Password string
	Host string
	Port int
	Name string
	Charset string
	Prefix string
}

var database *gorm.DB

func InitDatabase() error {
	var config Config
	err := viper.UnmarshalKey("database", &config)
	if err != nil {
		return err
	}

	// 创建数据库
	if err = createDatabaseIfNeed(config); err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
	}), &gorm.Config{
		// 禁用默认事务
		SkipDefaultTransaction: true,
		// 全局缓存预编译语句
		PrepareStmt: true,
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			// 表名前缀
			TablePrefix: config.Prefix,
			// 使用单数表名
			SingularTable: true,
		},
	})

	if err != nil {
		return err
	}

	database = db
	DB, err := db.DB()
	err = DB.Ping()

	if err != nil {
		return err
	}

	DB.SetMaxIdleConns(viper.GetInt("database.maxIdleConns"))
	DB.SetMaxOpenConns(viper.GetInt("database.maxOpenConns"))
	DB.SetConnMaxIdleTime(time.Hour)
	DB.SetConnMaxLifetime(5 * time.Minute)

	autoMigrate()

	return nil
}

// 自动迁移
func autoMigrate()  {
	database.AutoMigrate(&model.Account{})
}

// 获取数据库
func GetDatabase() *gorm.DB  {
	return database
}

// 创建数据库
func createDatabaseIfNeed (conf Config) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", conf.User, conf.Password, conf.Host, conf.Port))
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return err
	}
	var database string
	row := db.QueryRow(fmt.Sprintf("SHOW DATABASES LIKE '%s'", conf.Name))
	row.Scan(&database)
	if database == "" {
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARSET %s COLLATE %s_unicode_ci", conf.Name, conf.Charset, conf.Charset))
		if err != nil {
			return fmt.Errorf("create database err: %s", err)
		}
	}
	return nil
}