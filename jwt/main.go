package main

import (
	"FGA_Hacktiv8/jwt/database"
	"FGA_Hacktiv8/jwt/router"
)

func main() {
	database.StartDB()

	app := router.StartApp()

	app.Run(":8080")
}
