/*
 * © 2024 Alan Chan <alanchan@iis-corp.com>
 */
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	GitRoutes(router, "/git")

	log.Println("Repo Downloader starts to listen")
	router.Run("0.0.0.0:8088")
}
