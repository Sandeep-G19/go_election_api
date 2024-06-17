package routes

import (
	"database/sql"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitVotersRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/voters", func(c *gin.Context) { controllers.GetVoters(c, db) })
	r.GET("/voters/:id", func(c *gin.Context) { controllers.GetVoter(c, db) })
	r.POST("/voters", func(c *gin.Context) { controllers.CreateVoter(c, db) })
	r.PUT("/voters/:id", func(c *gin.Context) { controllers.UpdateVoter(c, db) })
	r.DELETE("/voters/:id", func(c *gin.Context) { controllers.DeleteVoter(c, db) })
}
