package rest

import (
	"github.com/DosyaKitarov/market-sniper/api"
	product "github.com/DosyaKitarov/market-sniper/internal/app/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	if c.Request.URL.Path != "/" {
		c.JSON(http.StatusNotFound, gin.H{"message": "404 Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Market Sniper"})
}

type RequestBody struct {
	Country string `json:"country"`
	Tld     string `json:"tld"`
}

func GetInfo(c *gin.Context) {
	asin := c.Param("ASIN")

	if asin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ASIN is required"})
		return
	}

	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body is not valid"})
		return
	}

	if body.Country == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Country is required"})
		return
	}

	if body.Tld == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "TLD is required"})
		return
	}

	// Fetch product data
	productData, err := api.FetchProductData(c, asin, body.Country, body.Tld)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var p product.Product
	p.FormatProduct(asin, productData)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, p)
}
