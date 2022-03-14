package main

import (
	"log"

	"github.com/KnBrBz/snippet-swagger/api/auth"
	"github.com/KnBrBz/snippet-swagger/api/endpoint"
	"github.com/KnBrBz/snippet-swagger/api/middleware"
	_ "github.com/KnBrBz/snippet-swagger/docs"
	"github.com/KnBrBz/snippet-swagger/setup"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	stp := setup.New()
	r := gin.New()

	r.Use(middleware.CORS(stp))

	auth.Init(r)
	endpoint.Init(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(stp.Host()); err != nil {
		log.Panicf("Run error: %v", err)
	}
}
