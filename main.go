package main

import (
	"rsv-system-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	_ = r.Run(":8080")
}
