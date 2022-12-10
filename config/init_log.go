package config

import (
	"github.com/airren/echo-bio-backend/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetLogger() {
	// 1 用new自定义log日志
	// zap.New(xxx)
	// 2 zap.New需要接收一个core，core是zapcore.Core类型，zapcore.Core是一个interface类型，
	//   而zapcore.NewCore返回的ioCore刚好实现了这个接口类型的所有5个方法，那么NewCore也可以认为是core类型
	// 3 所以zap.N
	//ew(core)变成了zap.New(zapcore.NewCore)
	// 4 而zapcore.NewCore需要三个变量：Encoder, WriteSyncer, LevelEnabler,我们在创建NewCore时自定义这三个类型变量即可，其中：
	//         Encoder：编码器 (写入日志格式)
	//         WriteSyncer：指定日志写到哪里去
	//         LevelEnabler：日志打印级别
	// NewCore(enc Encoder, ws WriteSyncer, enab LevelEnabler)

	// 4.2 通过GetEncoder获取自定义的Encoder
	Encoder := GetEncoder()
	// 4.4 通过GetWriteSyncer获取自定义的WriteSyncer
	WriteSyncer := GetWriteSyncer()
	// 4.6 通过GetLevelEnabler获取自定义的LevelEnabler
	LevelEnabler := GetLevelEnabler()
	// 4.7 通过Encoder、WriteSyncer、LevelEnabler创建一个core
	newCore := zapcore.NewCore(Encoder, WriteSyncer, LevelEnabler)
	// 5 传递 newCore New一个logger
	//  zap.AddCaller(): 输出文件名和行号
	//  zap.Fields: 假如每条日志中需要携带公用的信息，可以在这里进行添加
	global.Logger = zap.New(newCore)
}

func GetEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,      // 默认换行符"\n"
			EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 日志等级序列为小写字符串，如:InfoLevel被序列化为 "info"
			EncodeTime:     zapcore.EpochTimeEncoder,       // 日志时间格式显示
			EncodeDuration: zapcore.SecondsDurationEncoder, // 时间序列化，Duration为经过的浮点秒数
			EncodeCaller:   zapcore.ShortCallerEncoder,     // 日志行号显示
		})
}

// GetWriteSyncer 自定义的WriteSyncer 4.3
func GetWriteSyncer() zapcore.WriteSyncer {
	file, _ := os.Create("./zap.log")
	return zapcore.AddSync(file)
}

// GetLevelEnabler 自定义的LevelEnabler 4.5
func GetLevelEnabler() zapcore.Level {
	return zapcore.InfoLevel // 只会打印出info及其以上级别的日志
}
