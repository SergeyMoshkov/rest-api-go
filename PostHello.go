package main

import "github.com/gin-gonic/gin"

func PostHello(c *gin.Context) {
	c.JSON(200, gin.H{
        "message": "Hello REST API - HTTP POST!",
    })
}