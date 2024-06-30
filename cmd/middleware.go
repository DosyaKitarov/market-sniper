package main

import (
	"github.com/DosyaKitarov/market-sniper/internal/pkg/env"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) CheckApiKey(c *gin.Context) {
	incomingApiKey := env.GetEnvVariable("INCOMING_API_KEY")
	apiKey := c.GetHeader("X-Api-Key")
	switch {
	case apiKey == "":
		app.logger.Error("API key is missing")
		app.clientError(c.Writer, http.StatusUnauthorized)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API key is missing"})
		return
	case apiKey != incomingApiKey:
		app.logger.Error("Invalid API key")
		app.clientError(c.Writer, http.StatusForbidden)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API key is missing"})
		return
	case apiKey == incomingApiKey:
		c.Next()
	}

}
