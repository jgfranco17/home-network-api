package logger

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/jgfranco17/home-network-api/core/pkg/context_settings"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLoggerCanBeCreated(t *testing.T) {
	examples := []struct {
		description    string
		context        context.Context
		expectedFields map[string]string
	}{
		{
			description:    "Nil context",
			context:        nil,
			expectedFields: make(map[string]string),
		},
		{
			description:    "Empty background context",
			context:        context.Background(),
			expectedFields: make(map[string]string),
		},
		{
			description:    "Background context with request id",
			context:        context.WithValue(context.Background(), context_settings.RequestId, "some-request-id"),
			expectedFields: map[string]string{context_settings.RequestId: "some-request-id"},
		},
		{
			description:    "Background context with origin",
			context:        context.WithValue(context.Background(), context_settings.Origin, "vertex-studio@1.2.3"),
			expectedFields: map[string]string{context_settings.Origin: "vertex-studio@1.2.3"},
		},
	}

	for _, example := range examples {
		t.Run(example.description, func(t *testing.T) {
			logger := FromContext(example.context)
			assert.NotNilf(t, logger, "the logger could not be created.")

			logger.Logger.SetLevel(logrus.DebugLevel)
			buf := bytes.Buffer{}
			logger.Logger.SetOutput(&buf)

			logger.Error("")

			loggedText := buf.String()

			for key, value := range example.expectedFields {
				expectedText := fmt.Sprintf("%s=%s", key, value)
				assert.Contains(t, loggedText, expectedText)

			}

		})
	}

}

func TestLevelCanBeChangedWithStringValue(t *testing.T) {
	SetLevel("DEBUG")
	assert.Equal(t, logrus.GetLevel(), logrus.DebugLevel)

	SetLevel("INFO")
	assert.Equal(t, logrus.GetLevel(), logrus.InfoLevel)

	SetLevel("debug")
	assert.Equal(t, logrus.GetLevel(), logrus.DebugLevel)

	SetLevel("WARN")
	assert.Equal(t, logrus.GetLevel(), logrus.WarnLevel)

	SetLevel("not-correct")
	assert.Equal(t, logrus.GetLevel(), logrus.WarnLevel)

}
