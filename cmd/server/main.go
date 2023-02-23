package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phankieuphu/go-lang-base/config"
	"github.com/phankieuphu/go-lang-base/internal/model"
)

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
	// http.HandleFunc("/hello", helloHandler)
	// log.Println("Listing for requests at http://localhost:8000/hello")
	// log.Fatal(http.ListenAndServe(":8000", nil))
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello,world!",
		})
	})
	r.Run()
}
