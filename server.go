package main

import (
	"github.com/Moonlight-Zhao/go-project-example/cotroller"
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
