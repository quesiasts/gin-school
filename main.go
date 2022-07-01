package main

import (
	"github.com/quesiasts/gin-school/database"
	"github.com/quesiasts/gin-school/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
