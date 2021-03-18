package app

import (
	"fmt"
	"log"

	. "vvvstore/internal/app/controller"
	"vvvstore/internal/app/middleware/auth"
	"vvvstore/internal/app/middleware/cors"
	"vvvstore/internal/app/middleware/ginzap"
	"vvvstore/internal/app/middleware/trace"
	"vvvstore/internal/pkg/casbin"
	"vvvstore/internal/pkg/database"
	"vvvstore/internal/pkg/logger"
	"vvvstore/internal/pkg/redis"
	"vvvstore/internal/pkg/server"
	"vvvstore/internal/pkg/translator"

	"github.com/spf13/viper"
)

// 应用初始化
func Initialize() {
	fmt.Println("cors", viper.GetBool("app.cors"))
	// 初始化日志记录器
	logger.InitLogger("logs/app.log", viper.GetString("app.mode"))
	defer logger.Sync()

	// 初始化数据库
	if err := database.InitDatabase(); err != nil {
		log.Fatal(err)
	}

	// 初始化Redis
	if err := redis.InitRedisClient(); err != nil {
		log.Fatal(err)
	}

	// 翻译器初始化
	if err := translator.InitTrans(viper.GetString("app.locale")); err != nil {
		log.Fatal(err)
	}

	// casbin初始化
	e, err := casbin.InitCasbin()
	if err != nil {
		log.Fatal(err)
	}

	//e.AddPolicy("dajun", "data1", "read") // 添加一条规则
	//e.AddPolicy("lizi", "data2", "write")
	//e.AddPolicy("lizi", "data2", "read")


	ok, _ := casbin.Enforce(e, "root", "data1", "write")
	fmt.Println("权限检查:", ok)
	//e.SavePolicy() // 保存规则


	// 初始化定时任务管理器
	//cron.InitCronManger()
}

// 运行web服务
func Run()  {
	// 请求日志记录器
	logger := logger.NewAccessLogger()
	defer logger.Sync()

	// 服务初始化
	srv := server.Init()

	// 全局中间件
	srv.Use(
		ginzap.Ginzap(logger, "2006-01-02 15:04:05.000", false),
		ginzap.RecoveryWithZap(logger, true),
		cors.New(),
		trace.RequestId(),
	)

	// 分组1
	srv.Group("/api/v1", auth.JwtAuthorization())
	{
		srv.Router(NewIndexController())
	}

	// 分组2
	srv.Group("/api/admin")
	{
		srv.Router(NewUserController())
	}

	// 监听端口
	srv.Listen(fmt.Sprintf("%s:%d",viper.GetString("app.host"), viper.GetInt("app.port")))
}






