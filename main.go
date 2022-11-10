package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello restapigo")
	})
	r.GET("/hello/:id", func(c *gin.Context) {
		id := c.Param("id")
        c.String(http.StatusOK, fmt.Sprintf("hello %s", id))
	})
	r.GET("/hello2", hello)
	fmt.Println("Server run and listening")
	r.Run(":8080")
	

}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "hello restapigo 2")
}