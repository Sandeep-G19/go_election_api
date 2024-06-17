package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetElections(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, date, description FROM elections")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	var elections []gin.H
	for rows.Next() {
		var id int
		var name, description string
		var date string
		if err := rows.Scan(&id, &name, &date, &description); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		elections = append(elections, gin.H{
			"id":          id,
			"name":        name,
			"date":        date,
			"description": description,
		})
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with rows"})
		return
	}

	log.Printf("Elections: %v", elections)
	c.JSON(http.StatusOK, elections)
}

func GetElection(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	row := db.QueryRow("SELECT id, name, date, description FROM elections WHERE id = ?", id)

	var election gin.H
	var name, description string
	var date string
	if err := row.Scan(&id, &name, &date, &description); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No election found with id: %v", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "No election found"})
		} else {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
		}
		return
	}
	election = gin.H{
		"id":          id,
		"name":        name,
		"date":        date,
		"description": description,
	}
	log.Printf("Election: %v", election)
	c.JSON(http.StatusOK, election)
}

func CreateElection(c *gin.Context, db *sql.DB) {
	var election struct {
		Name        string `json:"name"`
		Date        string `json:"date"`
		Description string `json:"description"`
	}
	if err := c.BindJSON(&election); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO elections (name, date, description) VALUES (?, ?, ?)", election.Name, election.Date, election.Description)
	if err != nil {
		log.Fatalf("Error inserting into database: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error getting last insert ID: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateElection(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var election struct {
		Name        string `json:"name"`
		Date        string `json:"date"`
		Description string `json:"description"`
	}
	if err := c.BindJSON(&election); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE elections SET name = ?, date = ?, description = ? WHERE id = ?", election.Name, election.Date, election.Description, id)
	if err != nil {
		log.Fatalf("Error updating database: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func DeleteElection(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM elections WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from database: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
