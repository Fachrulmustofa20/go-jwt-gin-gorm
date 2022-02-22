package main

import (
	"go-jwt/database"
	"go-jwt/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
