package main

import (
	"github.com/faizallmaullana/libraryManagement/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// initialize the cors middleware
var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

func main() {
	r := gin.Default()

	// connect to the database
	models.ConnectDatabase()
	r.Use(cors.New(corsConfig))

	// ROUTES

	// run the server
	r.Run(":8080")
}
