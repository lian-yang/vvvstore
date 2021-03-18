package server

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/semihalev/gin-stats"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"vvvstore/docs"
)



// 定义一个规范接口
type IController interface {
	// 这个传参就是脚手架主程
	Router(server *Server)
}

// 定义一个脚手架
type Server struct {
	*gin.Engine
	// 路由分组
	RouterGroup *gin.RouterGroup
}

// 初始化函数
func Init() *Server {
	// 设置swagger接口文档基本信息
	docs.SwaggerInfo.Title = viper.GetString("app.name")
	docs.SwaggerInfo.Description = viper.GetString("app.description")
	docs.SwaggerInfo.Version = viper.GetString("app.version")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// 设置启动模式
	gin.SetMode(viper.GetString("app.mode"))

	// 自定义日志格式
	/*gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%s\t%s\t%s\t(%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}*/

	// 输出日志到文件和console。
	//f, _ := os.Create("access.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 初始化gin
	engine := gin.New()

	// 设置表单缓存最大内存
	engine.MaxMultipartMemory = 100 << 20

	// 静态文件服务
	engine.StaticFS("/static", http.Dir("static"))
	engine.StaticFS("/uploads", http.Dir("uploads"))
	engine.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 使用中间件
	engine.Use(stats.RequestStats())

	// 路由匹配失败处理
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})

	// 请求指标监控
	engine.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	// API文档
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 作为Server的构造器
	s := &Server{Engine: engine}

	// 返回作为链式调用
	return s
}

// 监听函数
func (s *Server) Listen(addr string) {
	err := endless.ListenAndServe(addr, s)
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}

// 挂载路由
func (s *Server) Router(controllers ...IController) *Server {
	// 遍历所有的控制器，这里使用接口，就是为了将Router实例化
	for _, c := range controllers {
		c.Router(s)
	}
	return s
}

// 挂载全局中间件
func (s *Server) Use(middlewares ...gin.HandlerFunc) *Server {
	s.Engine.Use(middlewares...)
	return s
}

// 路由分组
func (s *Server) Group(relativePath string, middlewares ...gin.HandlerFunc) *Server {
	s.RouterGroup = s.Engine.Group(relativePath)
	s.RouterGroup.Use(middlewares...)
	return s
}