package main

import (
	"github.com/goddamnnoob/family-app-api/app"
	"github.com/goddamnnoob/family-app-api/logger"
)

func main() {
	logger.Info("Starting app")
	app.Start()
}
