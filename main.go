package main

import (
	"groupie-tracker/internal/elk"
	"groupie-tracker/internal/server"
)

func main() {
	logger := elk.GetLogger()
	logger.Info("Server launching...")
	server.Runserver()
}
