package starters

import (
	"webapp/core/config"
	"webapp/core/config/viper"
	"webapp/core/logging"
	"webapp/core/logging/logrus"
	"webapp/core/validation"
	"webapp/core/validation/goplayground"
)

func setGlobalLogger() {
	logging.SetGlobal(logrus.New())
	logging.SetLevel(logging.DebugLevel)
	logging.SetFormatter(logging.TextFormat)
}

func setGlobalParser() {
	config.SetGlobal(viper.NewParserFromFile("webapp.yaml", "."))
}

func setGlobalValidator() {
	validation.SetGlobal(goplayground.New())
}

// Init initialize defaults values
func Init() {
	// Set Global Logger
	setGlobalLogger()
	// Set Global Parser
	setGlobalParser()
	// Set global Validator
	setGlobalValidator()
}
