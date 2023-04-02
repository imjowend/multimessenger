package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Hola Mundo!",
		})
	})

	r.Run() // Por defecto, escuchará en :8080
}

func run() error {
	a := gin.Default().RouterGroup
	fmt.Println(a)
	return nil
}
