package controllers

import (
	"fmt"
	"net/http"
	"organization-service/database"
	"organization-service/models"

	"github.com/gin-gonic/gin"
)

func AddMemberToOrganization(c *gin.Context) {
	var orgmember models.OrganizationMember
	if err := c.ShouldBindJSON(&orgmember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingMember models.OrganizationMember
	if err := database.DB.Where("organization_id = ? AND user_id = ?", orgmember.OrganizationID, orgmember.MemberID).First(&existingMember).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already a orgmember"})
		return
	}

	database.DB.Create(&orgmember)
	c.JSON(http.StatusCreated, orgmember)
}

func GetOrganizationMembers(c *gin.Context) {
	orgID := c.Param("id")
	var members []models.Member
	database.DB.Where("organization_id = ?", orgID).Find(&members)
	c.JSON(http.StatusOK, members)
}

func RemoveMemberFromOrganization(c *gin.Context) {
	orgID := c.Param("id")
	userID := c.Param("user_id")
	fmt.Println("orgID: ", orgID)
	fmt.Println("userID: ", userID)
	var orgmember models.Member
	if err := database.DB.Where("organization_id = ? AND user_id = ?", orgID, userID).First(&orgmember).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found in organization"})
		return
	}

	database.DB.Delete(&orgmember)
	c.JSON(http.StatusOK, gin.H{"message": "User removed from organization"})
}

func UpdateMemberRole(c *gin.Context) {
	orgID := c.Param("id")
	userID := c.Param("user_id")

	var orgmember models.OrganizationMember
	if err := database.DB.Where("organization_id = ? AND user_id = ?", orgID, userID).First(&orgmember).Error; err != nil {
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
	orgmember.Role = updateData.Role
	database.DB.Save(&orgmember)
	c.JSON(http.StatusOK, orgmember)
}
