package main

import (
	"organization-service/controllers"
	"organization-service/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	r := gin.Default()

	orgRoutes := r.Group("/organizations")
	{
		orgRoutes.POST("/", controllers.CreateOrganization)
		orgRoutes.GET("/", controllers.GetOrganizations)
		orgRoutes.GET("/:id", controllers.GetOrganizationByID)
		orgRoutes.PUT("/:id", controllers.UpdateOrganization)
		orgRoutes.DELETE("/:id", controllers.DeleteOrganization)

		orgRoutes.POST("/:id/members", controllers.AddMemberToOrganization)
		orgRoutes.GET("/:id/members", controllers.GetOrganizationMembers)
		orgRoutes.DELETE("/:id/members/:member_id", controllers.RemoveMemberFromOrganization)
		orgRoutes.PUT("/:id/members/:member_id", controllers.UpdateMemberRole)
	}
	membRoutes := r.Group("/members")
	{
		membRoutes.POST("/", controllers.CreateMember)
		membRoutes.GET("/", controllers.GetMembers)
		membRoutes.GET("/:id", controllers.GetMemberByID)
		membRoutes.PUT("/:id", controllers.UpdateMember)
		membRoutes.DELETE("/:id", controllers.DeleteMember)
	}

	r.Run(":8081")
}
