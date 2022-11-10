package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", GetHello)
	r.POST("/", PostHello)
	r.PUT("/", PutHello)
	r.DELETE("/", DeleteHello)
	r.Run(":8080")
	fmt.Println("server listening on")

}
