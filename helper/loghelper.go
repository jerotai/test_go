package helper


import (
	l4g "github.com/alecthomas/log4go"
)

type Log4go struct {
	logObject *l4g.Logger
}

// log config default
const format string  = "[%T %D] [%L] %M"

var HelperLog Log4go

func init() {
	log := &l4g.Logger{}
	logSetting := l4g.NewConsoleLogWriter()
	logSetting.SetFormat(format)
	log.AddFilter("console", l4g.DEBUG, logSetting)
	log.AddFilter("console", l4g.FINE, logSetting)
	log.AddFilter("console", l4g.TRACE, logSetting)
	log.AddFilter("console", l4g.INFO, logSetting)
	log.AddFilter("console", l4g.WARNING, logSetting)
	log.AddFilter("console", l4g.ERROR, logSetting)
	log.AddFilter("console", l4g.CRITICAL, logSetting)
	log.AddFilter("console", l4g.FINEST, logSetting)
	
	HelperLog = Log4go {
		logObject: log,
	}
}

/**
 * DEBUG TYPE
 */
func (l *Log4go) DebugLog(mess string) {
	l.logObject.Debug(mess)
}

/**
 * TRACE TYPE
 */
func (l *Log4go) TraceLog(mess string) {
	l.logObject.Trace(mess)
}

/**
 * FINE TYPE
 */
func (l *Log4go) FineLog(mess string) {
	l.logObject.Fine(mess)
}

/**
 * FINEST TYPE
 */
func (l *Log4go) FinestLog(mess string) {
	l.logObject.Finest(mess)
}

/**
 * CRITICAL TYPE
 */
func (l *Log4go) CriticalLog(mess string) {
	l.logObject.Critical(mess)
}

/**
 * INFO TYPE
 */
func (l *Log4go) InfoLog(mess string) {
	l.logObject.Info(mess)
}

/**
 * ERROR TYPE
 */
func (l *Log4go) ErrorLog(mess string) {
	l.logObject.Error(mess)
}

/**
 * WARING TYPE
 */
func (l *Log4go) WaringLog(mess string) {
	l.logObject.Warn(mess)
}

/**
 * 如果不是一直運行的程序,請加上這句話,否則主線程結束後,也不會輸出和log到日誌文件
 * defer helper.HelperLog.CloseLog()
 */
func (l *Log4go) CloseLog() {
	l.logObject.Close()
}