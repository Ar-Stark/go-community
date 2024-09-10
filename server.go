package main

import (
	"go-community/controller"
	"go-community/repository"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"os"
)

func main() {
	if err := repository.Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(http.StatusOK, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
