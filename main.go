package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// HTTP GET, POST, PUT, and DELETE requests
	r := gin.Default()
	r.GET("/", GetHello)
	r.POST("/", PostHello)
	r.PUT("/", PutHello)
	r.DELETE("/", DeleteHello)

	// group requests
	r1 := r.Group("/api") 
	{
        r1.GET("/hello", GetHello)
        r1.POST("/hello", PostHello)
        r1.PUT("/hello", PutHello)
		r1.DELETE("/hello", DeleteHello)
	}

	r.Run(":8080")
	fmt.Println("server listening on")

}
