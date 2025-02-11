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

func setupMemberRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/members", CreateMember)
	r.GET("/members", GetMembers)
	r.GET("/members/:id", GetMemberByID)
	r.PUT("/members/:id", UpdateMember)
	r.DELETE("/members/:id", DeleteMember)
	return r
}

func TestAddMember(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organization_members")
	defer database.TestDB.Exec("DELETE FROM members")
	defer database.TestDB.Exec("DELETE FROM organizations")

	r := setupMemberRouter()

	member := models.Member{Name: "Test Member", Email: "test@example.com", Password: "123456"}
	jsonValue, _ := json.Marshal(member)

	req, _ := http.NewRequest("POST", "/members", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetMembers(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM members")

	r := setupMemberRouter()

	req, _ := http.NewRequest("GET", "/members", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetMemberByID(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM members")

	member := models.Member{Name: "Test Member", Email: "test@example.com"}
	database.TestDB.Create(&member)

	r := setupMemberRouter()
	req, _ := http.NewRequest("GET", "/members/"+fmt.Sprint(member.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateMember(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM members")

	member := models.Member{Name: "Test Member", Email: "test@example.com", Password: "123456"}
	database.TestDB.Create(&member)

	r := setupMemberRouter()

	updatedMember := models.Member{Name: "Updated Member", Email: "updated@example.com", Password: "123456"}
	jsonValue, _ := json.Marshal(updatedMember)

	req, _ := http.NewRequest("PUT", "/members/"+fmt.Sprint(member.ID), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteMember(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM members")

	member := models.Member{Name: "Test Member", Email: "test@example.com"}
	database.TestDB.Create(&member)

	r := setupMemberRouter()

	req, _ := http.NewRequest("DELETE", "/members/"+fmt.Sprint(member.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
