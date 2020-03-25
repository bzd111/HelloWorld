package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainHandler(c *gin.Context) {
	htmlStr := `<html>
	<head></head>
	<body>
	<center>
	<h1>Hey Gophers!</h1>
	<h3>cat and dog picture</h3>
	<p><a href="/kitty">Cats</a></p>
	<p><a href="/pup">Dogs</a></p>
	</center>
	</body>
	</html>`

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(htmlStr))
}
