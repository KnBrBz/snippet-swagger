package main

import (
	"log"

	"github.com/KnBrBz/snippet-swagger/api/auth"
	"github.com/KnBrBz/snippet-swagger/api/endpoint"
	"github.com/KnBrBz/snippet-swagger/api/middleware"
	"github.com/KnBrBz/snippet-swagger/setup"
	"github.com/gin-gonic/gin"
)

func main() {
	stp := setup.New()
	r := gin.New()

	r.Use(middleware.CORS(stp))

	auth.Init(r)
	endpoint.Init(r)

	if err := r.Run(stp.Host()); err != nil {
		log.Panicf("Run error: %v", err)
	}
}
