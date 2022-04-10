package util

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() error {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		return err
	}
	return nil
}
