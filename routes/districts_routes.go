package routes

import (
	"database/sql"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitDistrictsRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/districts", func(c *gin.Context) { controllers.GetDistricts(c, db) })
	r.GET("/districts/:id", func(c *gin.Context) { controllers.GetDistrict(c, db) })
	r.POST("/districts", func(c *gin.Context) { controllers.CreateDistrict(c, db) })
	r.PUT("/districts/:id", func(c *gin.Context) { controllers.UpdateDistrict(c, db) })
	r.DELETE("/districts/:id", func(c *gin.Context) { controllers.DeleteDistrict(c, db) })
}
