/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logging

import (
	//
	"fmt"

	//
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	//
	"cluster-autoscaler-status-exporter/api/v1alpha1"
)

func ConfigureLogger(app *v1alpha1.ApplicationSpec) (err error) {
	// Define logger level. Default is Info
	level := zapcore.InfoLevel
	if app.Config.Logging.Level != "" {
		switch app.Config.Logging.Level {
		case "debug":
			level = zapcore.DebugLevel
		case "info":
			level = zapcore.InfoLevel
		case "warn":
			level = zapcore.WarnLevel
		case "error":
			level = zapcore.ErrorLevel
		case "dpanic":
			level = zapcore.DPanicLevel
		case "panic":
			level = zapcore.PanicLevel
		case "fatal":
			level = zapcore.FatalLevel
		default:
			level = zapcore.InfoLevel
		}
	}

	// Create a new logger
	logConfig := zap.Config{
		Encoding:         app.Config.Logging.Format,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	// Set timestamp format
	logConfig.EncoderConfig.TimeKey = "time"
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Build the logger
	logger, err := logConfig.Build()
	if err != nil {
		return fmt.Errorf("error creating logger: %v", err)
	}
	defer logger.Sync()

	// Enable with caller
	app.Logger = logger.WithOptions(zap.AddCaller())

	return err
}
