package initialize

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"os"
	"server/global"
	"server/task"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (z *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	z.logger.Info(msg, zap.Any("KeysAndValues", keysAndValues))
}
func (z *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	z.logger.Error(msg, zap.Error(err), zap.Any("KeysAndValues", keysAndValues))
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{
		logger: global.Log,
	}
}

func InitCron() {
	c := cron.New(cron.WithLogger(NewZapLogger()))
	err := task.RegisterScheduleTask(c)
	if err != nil {
		global.Log.Error("Failed to register schedule task")
		os.Exit(1)
	}
	c.Start()
}
