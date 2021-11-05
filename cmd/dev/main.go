package main

import (
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/app"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/logger"
)

func main(){
	logger.Info("Starting the application...")
	app.Start()
}

