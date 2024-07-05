package rest

import (
	"github.com/DosyaKitarov/market-sniper/api/email"
	product "github.com/DosyaKitarov/market-sniper/internal/app/product"
	"github.com/DosyaKitarov/market-sniper/internal/pkg/logger"
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

func (s *Services) GetProducts(c *gin.Context) {
	logger.Info(c, "GetProducts request received")

	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(c, "GetProducts.ShouldBindJSON() error: ", err)
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
	productArray, err := s.productService.InsertProducts(c, body.Asins, body.Country, body.Tld)
	if err != nil {
		logger.Error(c, "GetProducts.InsertProducts() error: ", err)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "processed asins": productArray})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, productArray)
}

func (s *Services) GetCsv(c *gin.Context) error {
	logger.Info(c, "GetCsv request received")

	productArray, err := s.productService.GetAllProducts(c)
	if err != nil {
		logger.Error(c, "GetCsv.GetAllProducts() error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return err
	}

	csvData, err := s.productService.ProductToCsv(c, productArray)
	if err != nil {
		logger.Error(c, "GetCsv.ProductToCsv() error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return err
	}

	err = email.SendCSVViaGmail(csvData)
	if err != nil {
		logger.Error(c, "GetCsv.SendCSVViaGmail() error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "CSV sent via email"})
	return nil
}
