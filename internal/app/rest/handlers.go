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
	Asins   []string `json:"asins"`
	Country string   `json:"country"`
	Tld     string   `json:"tld"`
}

func GetInfo(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body is not valid"})
		return
	}

	if body.Asins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "At least 1 ASIN is required"})
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

	var productArray []product.Product

	// Fetch product data
	for _, asin := range body.Asins {
		var p product.Product

		json, err := api.FetchProductData(c, asin, body.Country, body.Tld)
		if err != nil {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "processed asins": productArray})
			return
		}

		p.FormatProduct(asin, json)
		productArray = append(productArray, p)
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, productArray)
}
