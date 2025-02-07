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

func setupOrgMemberRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/organizations/:id/members", AddMemberToOrganization)
	r.GET("/organizations/:id/members", GetOrganizationMembers)
	r.DELETE("/organizations/:id/members/:member_id", RemoveMemberFromOrganization)
	r.PUT("/organizations/:id/members/:member_id", UpdateMemberRole)
	return r
}

func TestAddMemberToOrganization(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organization_members")
	defer database.TestDB.Exec("DELETE FROM members")
	defer database.TestDB.Exec("DELETE FROM organizations")

	// Gerekli kayıtları ekleyin
	member := models.Member{Name: "Test Member", Email: "test@example.com"}
	database.TestDB.Create(&member)
	org := models.Organization{Name: "Test Organization", Description: "Test Description"}
	database.TestDB.Create(&org)

	r := setupOrgMemberRouter()

	orgMember := models.OrganizationMember{OrganizationID: org.ID, MemberID: member.ID, Role: "member"}
	jsonValue, _ := json.Marshal(orgMember)

	req, _ := http.NewRequest("POST", fmt.Sprintf("/organizations/%d/members", org.ID), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetOrganizationMembers(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organization_members")
	defer database.TestDB.Exec("DELETE FROM members")
	defer database.TestDB.Exec("DELETE FROM organizations")

	// Gerekli kayıtları ekleyin
	member := models.Member{Name: "Test Member", Email: "test@example.com"}
	database.TestDB.Create(&member)
	org := models.Organization{Name: "Test Organization", Description: "Test Description"}
	database.TestDB.Create(&org)
	orgMember := models.OrganizationMember{OrganizationID: org.ID, MemberID: member.ID, Role: "member"}
	database.TestDB.Create(&orgMember)

	r := setupOrgMemberRouter()

	req, _ := http.NewRequest("GET", fmt.Sprintf("/organizations/%d/members", org.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRemoveMemberFromOrganization(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organization_members")
	defer database.TestDB.Exec("DELETE FROM members")
	defer database.TestDB.Exec("DELETE FROM organizations")

	// Gerekli kayıtları ekleyin
	member := models.Member{Name: "Test Member", Email: "test@example.com"}
	database.TestDB.Create(&member)
	org := models.Organization{Name: "Test Organization", Description: "Test Description"}
	database.TestDB.Create(&org)
	orgMember := models.OrganizationMember{OrganizationID: org.ID, MemberID: member.ID, Role: "member"}
	database.TestDB.Create(&orgMember)

	r := setupOrgMemberRouter()

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/organizations/%d/members/%d", org.ID, member.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateMemberRole(t *testing.T) {
	database.ConnectTestDB()
	defer database.TestDB.Exec("DELETE FROM organization_members")
	defer database.TestDB.Exec("DELETE FROM members")
	defer database.TestDB.Exec("DELETE FROM organizations")

	// Gerekli kayıtları ekleyin
	member := models.Member{Name: "Test Member", Email: "test@example.com"}
	database.TestDB.Create(&member)
	org := models.Organization{Name: "Test Organization", Description: "Test Description"}
	database.TestDB.Create(&org)
	orgMember := models.OrganizationMember{OrganizationID: org.ID, MemberID: member.ID, Role: "member"}
	database.TestDB.Create(&orgMember)

	r := setupOrgMemberRouter()

	updatedRole := struct {
		Role string `json:"role"`
	}{Role: "admin"}
	jsonValue, _ := json.Marshal(updatedRole)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/organizations/%d/members/%d", org.ID, member.ID), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
