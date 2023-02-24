package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phankieuphu/go-lang-base/config"
	"github.com/phankieuphu/go-lang-base/internal/model"
	"github.com/phankieuphu/go-lang-base/internal/routers"
)

var r = gin.Default()

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
	}
	if configs.Environment == "dev" {
		fmt.Println("Hello this is dev")
	}
	model.ConnetDatabase()
	fmt.Println("Hello, world!")
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Helgolo,world!",
		})
	})
	getRouters()
	r.Run()
}

func getRouters() {
	v1 := r.Group("/api/v1")
	routers.VoucherRouter(v1)
}
