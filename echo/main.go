package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.POST("/v1/alerts", func(ctx *gin.Context) {
		data, _ := ioutil.ReadAll(ctx.Request.Body)
		fmt.Println(string(data))
	})

	r.Run(":9094")
}
