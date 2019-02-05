package starters

import (
	"context"
	"webapp/core/config"
	"webapp/core/config/viper"
	"webapp/core/logging"
	"webapp/core/logging/logrus"
)

// Component interface for components to register
type Component interface {
	Init(ctx context.Context)
	Close(ctx context.Context)
}

func setGlobalLogger(_ context.Context) {
	logging.SetGlobal(logrus.New())
	logging.SetLevel(logging.DebugLevel)
	logging.SetFormatter(logging.TextFormat)
}

func setGlobalParser(_ context.Context) {
	config.SetGlobal(viper.NewParserFromFile("webapp.yaml", "."))
}

// Init initialize defaults values
func Init(ctx context.Context) {
	// Set Global Logger
	setGlobalLogger(ctx)
	// Set Global Parser
	setGlobalParser(ctx)

	for _, c := range Components() {
		c.Init(ctx)
	}
}

// Shutdown initialize defaults values
func Shutdown(ctx context.Context) {
	for _, c := range Components() {
		c.Close(ctx)
	}
}
