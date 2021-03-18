package job

import (
	"go.uber.org/zap"
	"time"
	"vvvstore/internal/pkg/logger"
)

var log *zap.Logger = logger.NewCronLogger()

type Test struct {

}

func (t Test) Spec() string {
	return "@every 3s" // Every 3 seconds
}

func (t Test) Run() {
	log.Info("test crontab run", zap.String("at", time.Now().Format("2006-01-02 15:04:05")))
}
