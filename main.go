package main

import (
	"github.com/mmmajder/devops-booking-service/startup"
	cfg "github.com/mmmajder/devops-booking-service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
