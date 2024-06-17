package routes

import (
	"database/sql"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitCampaignsRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/campaigns", func(c *gin.Context) { controllers.GetCampaigns(c, db) })
	r.GET("/campaigns/:id", func(c *gin.Context) { controllers.GetCampaign(c, db) })
	r.POST("/campaigns", func(c *gin.Context) { controllers.CreateCampaign(c, db) })
	r.PUT("/campaigns/:id", func(c *gin.Context) { controllers.UpdateCampaign(c, db) })
	r.DELETE("/campaigns/:id", func(c *gin.Context) { controllers.DeleteCampaign(c, db) })
}
