package main

import (
	"github.com/gin-gonic/gin"
	"go-web/src/router"

	"log"
	"net/http"
)

func main() {

	// Create the Gin engine.
	engine := gin.New()

	// gin middlewares
	var middlewares []gin.HandlerFunc

	// Routes.
	router.Load(
		// Cores.
		engine,

		// Middlewares.
		middlewares...,
	)

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", engine).Error())
}
