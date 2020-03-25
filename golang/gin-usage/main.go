package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// r.GET("/", mainHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Static("/img", "./img")
	r.GET("/", mainHandler)
	r.GET("/kitty", kittyHandler)
	r.GET("/pup", pupHandler)
	r.Run(":8000")
}
