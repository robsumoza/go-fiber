package main

import (
	"github.com/robsumoza/go-fiber/app"
	_ "github.com/robsumoza/go-fiber/docs"
)

// @title go-fiber
// @version 0.0.1
// @description An example template of a Golang simple backend API using Fiber
// @contact.name Roberto Sumoza
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	// setup and run app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
