package main

import "github.com/gin-gonic/gin"

func GetHello(c *gin.Context) {
	c.JSON(200, gin.H{
        "message": "Hello REST API - HTTP GET!",
    })
}