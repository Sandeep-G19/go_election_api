package routes

import (
	"database/sql"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitCandidatesRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/candidates", func(c *gin.Context) { controllers.GetCandidates(c, db) })
	r.GET("/candidates/:id", func(c *gin.Context) { controllers.GetCandidate(c, db) })
	r.POST("/candidates", func(c *gin.Context) { controllers.CreateCandidate(c, db) })
	r.PUT("/candidates/:id", func(c *gin.Context) { controllers.UpdateCandidate(c, db) })
	r.DELETE("/candidates/:id", func(c *gin.Context) { controllers.DeleteCandidate(c, db) })
}
