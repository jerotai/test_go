package helper

import (
	"testing"
	
	"Stingray/helper"
)

// 測試 Log4Go
func TestLog4go(t *testing.T) {
	helper.HelperLog.TraceLog("trace test")
	helper.HelperLog.InfoLog("info test")
	helper.HelperLog.ErrorLog("error test")
	helper.HelperLog.WaringLog("warn test")
	helper.HelperLog.CriticalLog("crit test")
	helper.HelperLog.FineLog("fine test")
	helper.HelperLog.FinestLog("finest test")
	helper.HelperLog.DebugLog("debug test")
	defer helper.HelperLog.CloseLog()
}