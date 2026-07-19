package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		h, err := os.Hostname()
		if err != nil {
			h = "unknown-host"
		}
		c.String(http.StatusOK, fmt.Sprintf("hello from %s", h))
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	_ = r.Run(":8080")
}
