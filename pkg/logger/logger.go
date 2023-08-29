package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/czjge/gohub/pkg/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 全局 Logger 对象
var Logger *zap.Logger

// 日志初始化
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {

	// 获取日志写入介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	// 设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志级别设置有误")
	}

	// 初始化 core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// 初始化 Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)
}

// 设置日志存储格式
func getEncoder() zapcore.Encoder {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 哪里调用的代码，如：paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,   // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	if app.IsLocal() {
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

// 自定义日志时间输出格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// 记录日志的介质：终端输出和文件
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {

	if logType == "daily" {
		logname := time.Now().Format("2006-01-02") + ".log"
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackup,
		MaxAge:     maxAge,   //days
		Compress:   compress, // disabled by default
	}

	// log rotation is only activated by log file size,
	// if you want to activate it by day:
	// go func() {
	// 	for {
	// 		nowTime := time.Now()
	// 		nowTimeStr := nowTime.Format("2006-01-02")
	// 		t2, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.Local)
	// 		next := t2.AddDate(0, 0, 1)
	// 		after := next.UnixNano() - nowTime.UnixNano() - 1
	// 		<-time.After(time.Duration(after) * time.Nanosecond)
	// 		lumberJackLogger.Rotate()
	// 	}
	// }()

	if app.IsLocal() {
		// 终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		// 只记录文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

// 调试专用，不中断应用
// eg：logger.Dump(user.User{Name:"test"}, "用户信息")
func Dump(value any, msg ...string) {
	valueString := jsonString(value)
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

// 当 err != nil 时候记录 error 等级的错误
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试类日志
// eg：logger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// 不会退出程序
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// 写完日志后调用 os.Exit(1) 退出程序。错误级别和 Error 一样
func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

// 记录一条字符串类型的 debug 日志
// eg：logger.DebugString("SMS", "短信内容", string(result.RawResponse))
func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

// 记录对象类型的 debug 日志
// eg：logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, value any) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(value any) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
