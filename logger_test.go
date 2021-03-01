package logger

import (
	"time"
    "testing"	
)

func TestLogger(t *testing.T) {
	defer Destory()

	Error("test_Error")
	Errorf("test_Errorf|%s,%s","成功",time.Unix(0, time.Now().UnixNano()).Format("2006-01-02 15:04:05.000000000"))
	Warning("test_Warning")
	Warningf("test_Warningf|%s,%s","成功",time.Unix(0, time.Now().UnixNano()).Format("2006-01-02 15:04:05.000000000"))
    Info("test_Info")
	Infof("test_Infof|%s,%s","成功",time.Unix(0, time.Now().UnixNano()).Format("2006-01-02 15:04:05.000000000"))
	Debug("test_Debug")
	Debugf("test_Debugf|%s,%s","成功",time.Unix(0, time.Now().UnixNano()).Format("2006-01-02 15:04:05.000000000"))
}