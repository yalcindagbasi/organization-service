package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"organization-service/database"
	"organization-service/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	orgRoutes := r.Group("/organizations")
	{
		orgRoutes.POST("/", CreateOrganization)
		orgRoutes.GET("/", GetOrganizations)
		orgRoutes.GET("/:id", GetOrganizationByID)
		orgRoutes.PUT("/:id", UpdateOrganization)
		orgRoutes.DELETE("/:id", DeleteOrganization)
	}
	return r
}

func TestCreateOrganization(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organizations")

	router := setupRouter()

	org := models.Organization{Name: "Test Organizasyonu", Description: "Test Açıklaması"}
	jsonValue, _ := json.Marshal(org)

	req, _ := http.NewRequest("POST", "/organizations/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetOrganizations(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organizations")
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/organizations/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetOrganizationByID(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organizations")

	router := setupRouter()

	org := models.Organization{Name: "Test Organizasyonu", Description: "Test Açıklaması"}
	database.TestDB.Create(&org)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/organizations/%d", org.ID), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateOrganization(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organizations")
	router := setupRouter()

	org := models.Organization{Name: "Test Organizasyonu", Description: "Test Açıklaması"}
	database.TestDB.Create(&org)

	updateOrg := models.Organization{Name: "Updated Name", Description: "Updated Desc"}
	jsonValue, _ := json.Marshal(updateOrg)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/organizations/%d", org.ID), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteOrganization(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organizations")

	router := setupRouter()

	org := models.Organization{Name: "Test Organizasyonu", Description: "Test Açıklaması"}
	database.DB.Create(&org)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/organizations/%d", org.ID), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
