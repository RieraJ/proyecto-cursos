package main

import (
	"backend/app"
	"backend/clients"
	"backend/initializers"
)

func main() {
	initializers.LoadEnvVariables()
	clients.ConnectDb()
	app.StartRoute()
}
