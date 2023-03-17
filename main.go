package main

import (
	"github.com/gin-gonic/gin"
	"ginProject/routes"
)

func main() {
	r := gin.Default()

	// initialize the routes
	routes.InitUserRoutes(r)

	r.Run(":8080")
}
