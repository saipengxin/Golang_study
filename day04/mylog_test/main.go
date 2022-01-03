package main

import "github.com/saipengxin/study/day04/mylog"

func main() {
	logger := mylog.NewLogger()
	logger.Info("Info等级的日志")
	logger.Trace("Trace等级的日志")
	logger.Info("Info等级的日志")
	logger.Warning("Warning等级的日志")
	logger.Error("Error等级的日志")
	logger.Fatal("Fatal等级的日志")
}
