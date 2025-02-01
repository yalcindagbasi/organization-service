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
	}

	r.Run(":8081")
}
