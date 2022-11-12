package main

import (
	"fmt"
	"strconv"

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

	// Handing path requests
	r.GET("product/:id", GetProductByID)
	r.GET("profile/:username", ShowProfile)
	r.GET("compute/:num1/add/:num2", Compute)

	r.Run(":8080")
	fmt.Println("server listening on")

}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	c.JSON(200, gin.H{
		"id":      id,
		"product": "Milk",
	})
}

func ShowProfile(c *gin.Context) {
	username := c.Param("username")
	fmt.Println(username)
	c.JSON(200, gin.H{
		"username": username,
		"family":   "Moshkov",
	})
}

func Compute(c *gin.Context) {
	num1, _ := strconv.Atoi(c.Param("num1"))
	num2, _ := strconv.Atoi(c.Param("num2"))
	fmt.Println(num1, num2)
	sum := num1 + num2
	c.JSON(200, gin.H{
		"num1": num1,
		"num2": num2,
		"sum":  sum,
	})
}
