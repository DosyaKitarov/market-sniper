package main

import (
	rest "github.com/DosyaKitarov/market-sniper/internal/app/rest"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	router := gin.Default()

	// Custom 404 handler
	router.NoRoute(func(c *gin.Context) {
		app.notFound(c.Writer)
	})

	// Middleware
	recoveryMiddleware := gin.Recovery()
	loggingMiddleware := gin.Logger()
	apiCheckMiddleware := app.CheckApiKey

	// Apply middleware
	router.Use(recoveryMiddleware, loggingMiddleware, apiCheckMiddleware)

	// Define routes
	router.GET("/", func(c *gin.Context) {
		rest.Home(c)
	})

	router.GET("/api/v1/getInfo", func(c *gin.Context) {
		rest.GetInfo(c)
	})

	return router
}
