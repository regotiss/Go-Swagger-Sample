package main

import (
	"sample-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	r.GET("/hello", displayHello)
	addSwaggerInfo()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

func addSwaggerInfo() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample hello server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "hello.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

//Hello godoc
//@Tags hello
//@Description simple hello message
//@Accept json
//@Produce json
//@Success 200 {object} Message
//@Router /hello [get]
func displayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

type Message struct {
	Msg string `json:"message"`
}
