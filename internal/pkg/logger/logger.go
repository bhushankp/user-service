package logger

import "go.uber.org/zap"

var Log *zap.Logger

func InitLogger() {

	l, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger:" + err.Error())
	}
	Log = l

	Log.Info("Logger Initialize")

}
