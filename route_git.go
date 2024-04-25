/*
 * © 2022 Alan Chan <alanchan@iis-corp.com>
 */
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
)

func GitRoutes(router *gin.Engine, apipath string) {
	router.OPTIONS(apipath+"/download", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Methods", "POST")
		c.Status(http.StatusOK)
	})

	router.POST(apipath, downloadRepo)
}

func downloadRepo(c *gin.Context) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	_, err = git.PlainClone(data["path"].(string), false, &git.CloneOptions{
		URL:      data["url"].(string),
		Progress: os.Stdout,
	})
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}
