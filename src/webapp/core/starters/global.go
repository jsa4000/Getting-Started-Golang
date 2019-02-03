package starters

import (
	"context"
	"webapp/core/config"
	"webapp/core/config/viper"
	"webapp/core/logging"
	"webapp/core/logging/logrus"
	"webapp/core/net/http"

	router "webapp/core/net/http/gorillamux"
	"webapp/core/storage/mongo"
	"webapp/core/validation"
	"webapp/core/validation/goplayground"
)

func setGlobalLogger(_ context.Context) {
	logging.SetGlobal(logrus.New())
	logging.SetLevel(logging.DebugLevel)
	logging.SetFormatter(logging.TextFormat)
}

func setGlobalParser(_ context.Context) {
	config.SetGlobal(viper.NewParserFromFile("webapp.yaml", "."))
}

func setGlobalValidator(_ context.Context) {
	validation.SetGlobal(goplayground.New())
}

func setGlobalRouter(_ context.Context) {
	http.SetGlobal(router.New())
}

func setGlobalMongo(ctx context.Context) {
	mongo.SetGlobal(mongo.New())
	mongo.Connect(ctx, "mongodb://root:root@dockerhost:27017/admin")
}

func unsetGlobalMongo(ctx context.Context) {
	mongo.Disconnect(ctx)
}

// Init initialize defaults values
func Init(ctx context.Context) {
	// Set Global Logger
	setGlobalLogger(ctx)
	// Set Global Parser
	setGlobalParser(ctx)
	// Set global Validator
	setGlobalValidator(ctx)
	// Set global Router
	setGlobalRouter(ctx)
	// Set global MongoDB Storage
	setGlobalMongo(ctx)
}

// Shutdown initialize defaults values
func Shutdown(ctx context.Context) {
	// Set global MongoDB Storage
	unsetGlobalMongo(ctx)
}
