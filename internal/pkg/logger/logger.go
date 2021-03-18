package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// 只能输出结构化日志，性能高于 SugaredLogger
var Log *zap.Logger

// 初始化日志记录器
func InitLogger(filename string, mode string)  {
	encoder := getEncoder()
	writer := getWriter(filename)
	logLevel := zap.InfoLevel
	if mode == "debug" {
		logLevel = zap.DebugLevel
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writer),
		logLevel)
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	// zap.Fields(zap.String("app", "vvvstore")
}

// 请求日志记录器
func NewAccessLogger() *zap.Logger {
	encoder := getEncoder()
	writer := getWriter("logs/access.log")
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writer),
		zap.DebugLevel)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

// cron日志记录器
// 请求日志记录器
func NewCronLogger() *zap.Logger {
	encoder := getEncoder()
	writer := getWriter("logs/cron.log")
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writer),
		zap.DebugLevel)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

// 获取日志写入器
func getWriter(filename string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename, // 日志文件位置
		MaxSize: 512, // 日志文件最大大小（MB）
		MaxBackups: 10, // 保留日志文件最大数
		MaxAge: 30, // 保留日志文件最长天数
		LocalTime: true, // 是否使用本地时间
		Compress: false, // 是否压缩日志文件
	})
}

// 获取日志编码器
func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.MessageKey = "message"
	encodeConfig.TimeKey = "time"
	encodeConfig.CallerKey = "file"
	encodeConfig.StacktraceKey = "stack"
	encodeConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	} // 自定义时间格式
	return zapcore.NewJSONEncoder(encodeConfig)
}

// 同步缓冲的日志
func Sync()  {
	Log.Sync()
}

func Debug(msg string, fields ...zap.Field)  {
	Log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field)  {
	Log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field)  {
	Log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field)  {
	Log.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field)  {
	Log.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field)  {
	Log.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field)  {
	Log.Panic(msg, fields...)
}




