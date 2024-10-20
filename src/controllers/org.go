package controllers

import (
	"net/http"

	"github.com/asahi-zip/api/models"

	"github.com/gin-gonic/gin"
)

func CreateOrg(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	org := models.Org{Name: input.Name}

	if err := models.DB.Create(&org).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create organization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "organization created successfully", "org": org})
}
