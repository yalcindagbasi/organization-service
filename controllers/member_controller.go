package controllers

import (
	"net/http"
	"organization-service/database"
	"organization-service/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": CustomErrorMessage(err)})
		return
	}
	var existingMember models.Member
	if err := database.DB.Where("email = ?", member.Email).First(&existingMember).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Member already exists"})
		return
	}

	if err := database.DB.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, member)
}

func GetMembers(c *gin.Context) {
	var members []models.Member
	database.DB.Find(&members)
	c.JSON(http.StatusOK, members)
}
func GetMemberByID(c *gin.Context) {
	id := c.Param("id")
	var member models.Member
	if err := database.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}
	c.JSON(http.StatusOK, member)
}
func UpdateMember(c *gin.Context) {
	id := c.Param("id")
	var member models.Member
	if err := database.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": CustomErrorMessage(err)})
		return
	}

	if err := database.DB.Save(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, member)
}
func DeleteMember(c *gin.Context) {
	id := c.Param("id")
	var member models.Member
	if err := database.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	if err := database.DB.Delete(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully"})
}

func CustomErrorMessage(err error) string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			errors = append(errors, e.Field()+" is required")
		case "email":
			errors = append(errors, e.Field()+" must be a valid email address")
		case "min":
			errors = append(errors, e.Field()+" must be at least "+e.Param()+" characters")
		case "max":
			errors = append(errors, e.Field()+" must be at most "+e.Param()+" characters")
		case "gte":
			errors = append(errors, e.Field()+" must be greater than or equal to "+e.Param())
		case "lte":
			errors = append(errors, e.Field()+" must be less than or equal to "+e.Param())
		}
	}
	return errors[0]
}
