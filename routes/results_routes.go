package routes

import (
	"database/sql"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitResultsRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/results", func(c *gin.Context) { controllers.GetResults(c, db) })
	r.GET("/results/:id", func(c *gin.Context) { controllers.GetResult(c, db) })
	r.POST("/results", func(c *gin.Context) { controllers.CreateResult(c, db) })
	r.PUT("/results/:id", func(c *gin.Context) { controllers.UpdateResult(c, db) })
	r.DELETE("/results/:id", func(c *gin.Context) { controllers.DeleteResult(c, db) })
}
