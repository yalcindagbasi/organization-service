package controllers

import (
	"net/http"
	"organization-service/database"
	"organization-service/models"

	"github.com/gin-gonic/gin"
)

func CreateOrganization(c *gin.Context) {
	var organization models.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": CustomErrorMessage(err)})
		return
	}

	database.DB.Create(&organization)
	c.JSON(http.StatusCreated, organization)
}

func GetOrganizations(c *gin.Context) {
	var organizations []models.Organization
	database.DB.Find(&organizations)
	c.JSON(http.StatusOK, organizations)
}

func GetOrganizationByID(c *gin.Context) {
	id := c.Param("id")
	var organization models.Organization
	if err := database.DB.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}
	c.JSON(http.StatusOK, organization)
}

func UpdateOrganization(c *gin.Context) {
	id := c.Param("id")
	var organization models.Organization
	if err := database.DB.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": CustomErrorMessage(err)})
		return
	}
	database.DB.Save(&organization)
	c.JSON(http.StatusOK, organization)
}

func DeleteOrganization(c *gin.Context) {
	id := c.Param("id")
	var organization models.Organization
	if err := database.DB.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	database.DB.Delete(&organization)
	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}
