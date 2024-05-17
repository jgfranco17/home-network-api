package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

func GetLogger(ctx context.Context) *logrus.Entry {
	logger := logrus.New()
	logLevel := logrus.InfoLevel
	logger.SetLevel(logLevel)
	return logger.WithFields(logrus.Fields{
		"Section": ctx.Value("section"),
	})
}
