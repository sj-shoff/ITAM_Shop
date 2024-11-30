package controllers

import (
	config "myapp/internal/data_base"
	entity "myapp/internal/structures"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductFilter(c *gin.Context) {
	var filterParams entity.FilterParams
	if err := c.ShouldBindJSON(&filterParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	var products []entity.Product
	query := config.DB

	if filterParams.Category != "" {
		query = query.Where("category = ?", filterParams.Category)
	}

	query = query.Where("price >= ? AND price <= ?", filterParams.MinPrice, filterParams.MaxPrice)

	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return

	}

	c.JSON(http.StatusOK, products)

}
