package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

// 初始化配置文件
func (c *Config) init() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 设置配置文件目录
		viper.SetConfigName("config") // 解析默认的配置文件
	}

	viper.SetConfigType("toml") // 设置配置文件格式为toml
	viper.AutomaticEnv() // 读取匹配的环境变量
	viper.SetEnvPrefix("VVVSTORE") // 读取环境变量的前缀为VVVSTORE
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // 重写Env定界符

	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

// 监听配置变化
func (c *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file changed: %s", e.Name)
	})
}

// 重写配置文件
func (c *Config) Write() error {
	return viper.SafeWriteConfig()
}

// 初始化配置
func InitConfig(name string) error {
	c := Config {
		Name: name,
	}
	
	// 初始化配置文件
	if err := c.init(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watch()

	return nil
}