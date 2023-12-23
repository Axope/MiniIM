package log

import (
	"testing"

	"go.uber.org/zap"
)

func TestSimpleHttpGet(t *testing.T) {
	InitLogger("DEBUG", "logs/test.log")
	defer Logger.Sync()

	Logger.Debug("this is debug...")
	Logger.Info("this is info...")
	Logger.Warn("this is warn...")
	Logger.Error("this is error...", zap.String("application", "MiniIM"))
}
