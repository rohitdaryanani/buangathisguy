package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// init with default middlewares i.e: logger, recovery etc...
	r := gin.Default()

	// static assets
	r.Static("/js", "public/js")
	r.Static("/css", "public/css")
	r.Static("/images", "public/images")

	// routes
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It Worked!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// serve
	r.Run(fmt.Sprintf(":%s", port))
}
