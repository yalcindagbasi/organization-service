package controllers

import (
	"fmt"
	"net/http"
	"organization-service/database"
	"organization-service/models"

	"github.com/gin-gonic/gin"
)

func AddMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingMember models.Member
	if err := database.DB.Where("organization_id = ? AND user_id = ?", member.OrganizationID, member.UserID).First(&existingMember).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already a member"})
		return
	}

	database.DB.Create(&member)
	c.JSON(http.StatusCreated, member)
}

func GetOrganizationMembers(c *gin.Context) {
	orgID := c.Param("id")
	var members []models.Member
	database.DB.Where("organization_id = ?", orgID).Find(&members)
	c.JSON(http.StatusOK, members)
}

func RemoveMember(c *gin.Context) {
	orgID := c.Param("id")
	userID := c.Param("user_id")
	fmt.Println("orgID: ", orgID)
	fmt.Println("userID: ", userID)
	var member models.Member
	if err := database.DB.Where("organization_id = ? AND user_id = ?", orgID, userID).First(&member).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found in organization"})
		return
	}

	database.DB.Delete(&member)
	c.JSON(http.StatusOK, gin.H{"message": "User removed from organization"})
}

func UpdateMemberRole(c *gin.Context) {
	orgID := c.Param("id")
	userID := c.Param("user_id")

	var member models.Member
	if err := database.DB.Where("organization_id = ? AND user_id = ?", orgID, userID).First(&member).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found in organization"})
		return
	}
	var updateData struct {
		Role string `json:"role"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	member.Role = updateData.Role
	database.DB.Save(&member)
	c.JSON(http.StatusOK, member)
}
