package main

import (
	"fmt"
	"log"
	"splitbill/database"
	"splitbill/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("---main called---")

	// initializing database connection
	database.ConnectDB()

	// closing database connection
	defer database.Postgres.Close()

	r := gin.Default()

	// myAppRouter := r.Group("app")

	router.EndPoints(r)

	log.Fatal(r.Run())
}
