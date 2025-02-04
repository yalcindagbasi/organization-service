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

		orgRoutes.POST("/:id/members", controllers.AddMember)
		orgRoutes.GET("/:id/members", controllers.GetOrganizationMembers)
		orgRoutes.DELETE("/:id/members/:user_id", controllers.RemoveMember)
		orgRoutes.PUT("/:id/members/:user_id", controllers.UpdateMemberRole)
	}
	r.Run(":8081")
}
