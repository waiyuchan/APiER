package gorm_v2

import (
	"apier/internal/global/variable"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strings"
	"time"
)

/** 自定义日志格式, 对 gorm 自带日志进行拦截重写 **/

// createCustomGormLog 创建自定义的 GORM 日志对象，用于拦截和重写 GORM 的日志输出
func createCustomGormLog(sqlType string, options ...Options) gormLog.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	// 从全局变量中获取配置信息
	logConf := gormLog.Config{
		SlowThreshold:             time.Second * variable.ConfigGormYaml.GetDuration("Gorm."+sqlType+"."+".SlowThreshold"),
		LogLevel:                  gormLog.Warn,
		Colorful:                  false,
		IgnoreRecordNotFoundError: variable.ConfigGormYaml.GetBool("Gorm." + sqlType + "." + ".IgnoreRecordNotFoundError"),
	}

	// 创建日志对象并应用配置选项
	log := &logger{
		Writer:       logOutPut{},
		Config:       logConf,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
	for _, val := range options {
		val.apply(log)
	}
	return log
}

// logOutPut 自定义的日志输出对象
type logOutPut struct{}

// Printf 根据格式化字符串和参数输出日志信息
func (l logOutPut) Printf(strFormat string, args ...interface{}) {
	logRes := fmt.Sprintf(strFormat, args...)
	logFlag := "gorm_v2 日志:"
	detailFlag := "详情："

	// 根据日志格式前缀判断日志级别，并输出相应级别的日志
	if strings.HasPrefix(strFormat, "[info]") || strings.HasPrefix(strFormat, "[traceStr]") {
		variable.ZapLog.Info(logFlag, zap.String(detailFlag, logRes))
	} else if strings.HasPrefix(strFormat, "[error]") || strings.HasPrefix(strFormat, "[traceErr]") {
		// 如果日志中包含 "record not found"，则不再将其视为错误记录到日志中
		if !strings.Contains(logRes, "record not found") {
			variable.ZapLog.Error(logFlag, zap.String(detailFlag, logRes))
		}
	} else if strings.HasPrefix(strFormat, "[warn]") || strings.HasPrefix(strFormat, "[traceWarn]") {
		variable.ZapLog.Warn(logFlag, zap.String(detailFlag, logRes))
	}
}

// Options 定义了用于设置日志格式的选项接口
type Options interface {
	apply(*logger)
}

// OptionFunc 实现了选项接口的函数类型
type OptionFunc func(log *logger)

// apply 将选项应用到日志对象上
func (f OptionFunc) apply(log *logger) {
	f(log)
}

// SetInfoStrFormat 设置 info 日志格式的选项函数
func SetInfoStrFormat(format string) Options {
	return OptionFunc(func(log *logger) {
		log.infoStr = format
	})
}

// SetWarnStrFormat 设置 warn 日志格式的选项函数
func SetWarnStrFormat(format string) Options {
	return OptionFunc(func(log *logger) {
		log.warnStr = format
	})
}

// SetErrStrFormat 设置 error 日志格式的选项函数
func SetErrStrFormat(format string) Options {
	return OptionFunc(func(log *logger) {
		log.errStr = format
	})
}

// SetTraceStrFormat 设置 trace 日志格式的选项函数
func SetTraceStrFormat(format string) Options {
	return OptionFunc(func(log *logger) {
		log.traceStr = format
	})
}

// SetTracWarnStrFormat 设置 traceWarn 日志格式的选项函数
func SetTracWarnStrFormat(format string) Options {
	return OptionFunc(func(log *logger) {
		log.traceWarnStr = format
	})
}

// SetTracErrStrFormat 设置 traceErr 日志格式的选项函数
func SetTracErrStrFormat(format string) Options {
	return OptionFunc(func(log *logger) {
		log.traceErrStr = format
	})
}

// logger 自定义的日志对象，实现了 gormLog.Interface 接口
type logger struct {
	gormLog.Writer
	gormLog.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode 设置日志模式
func (l *logger) LogMode(level gormLog.LogLevel) gormLog.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info 输出 info 级别日志
func (l logger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Info {
		l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn 输出 warn 级别日志
func (l logger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error 输出 error 级别日志
func (l logger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Error {
		l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace 输出 SQL 语句和执行情况的日志
func (l logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormLog.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormLog.Error && (!errors.Is(err, gormLog.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-1", sql)
		} else {
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLog.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-1", sql)
		} else {
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == gormLog.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-1", sql)
		} else {
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
