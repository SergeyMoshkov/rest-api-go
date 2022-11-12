package main

import "github.com/gin-gonic/gin"

func PutHello(c *gin.Context) {
	c.JSON(200, gin.H{
        "message": "Hello REST API - HTTP PUT!",
    })
}