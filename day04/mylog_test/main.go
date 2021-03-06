package main

import (
	"github.com/saipengxin/study/day04/mylog"
	"time"
)

func main() {
	for {
		logger := mylog.NewLogger("File", "Debug")
		var err = "123报错了"
		logger.Debug("Debug等级的日志,报错信息为:%v", err)
		logger.Trace("Trace等级的日志")
		logger.Info("Info等级的日志")
		logger.Warning("Warning等级的日志")
		logger.Error("Error等级的日志")
		logger.Fatal("Fatal等级的日志")
		time.Sleep(time.Second)
	}
}
