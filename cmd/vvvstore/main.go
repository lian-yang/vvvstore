package main

import (
	"fmt"
	"os"
	"vvvstore/cmd"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
const VERSION = "1.0.0"

// @title vvvstore商城系统
// @version 1.0.0
// @description 基于golang + uniapp 实现的跨端商城系统、支持H5 Android IOS 微信小程序等

// @contact.name yanglian
// @contact.url https://github.com/lian-yang
// @contact.email 395486566@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
