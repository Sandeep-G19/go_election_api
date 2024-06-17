package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCandidates(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, party_id FROM candidates")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	var candidates []gin.H
	for rows.Next() {
		var id int
		var name string
		var partyID int
		if err := rows.Scan(&id, &name, &partyID); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		candidates = append(candidates, gin.H{
			"id":       id,
			"name":     name,
			"party_id": partyID,
		})
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with rows"})
		return
	}

	log.Printf("Candidates: %v", candidates)
	c.JSON(http.StatusOK, candidates)
}

func GetCandidate(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	row := db.QueryRow("SELECT id, name, party_id FROM candidates WHERE id = ?", id)

	var candidate gin.H
	var name string
	var partyID int
	if err := row.Scan(&id, &name, &partyID); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No candidate found with id: %v", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "No candidate found"})
		} else {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
		}
		return
	}
	candidate = gin.H{
		"id":       id,
		"name":     name,
		"party_id": partyID,
	}
	log.Printf("Candidate: %v", candidate)
	c.JSON(http.StatusOK, candidate)
}

func CreateCandidate(c *gin.Context, db *sql.DB) {
	var candidate struct {
		Name    string `json:"name"`
		PartyID int    `json:"party_id"`
	}

	if err := c.BindJSON(&candidate); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := db.Exec("INSERT INTO candidates (name, party_id) VALUES (?, ?)", candidate.Name, candidate.PartyID)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	id, _ := result.LastInsertId()
	candidateWithID := gin.H{
		"id":       id,
		"name":     candidate.Name,
		"party_id": candidate.PartyID,
	}
	log.Printf("Candidate Created: %v", candidateWithID)
	c.JSON(http.StatusCreated, candidateWithID)
}

func UpdateCandidate(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var candidate struct {
		Name    string `json:"name"`
		PartyID int    `json:"party_id"`
	}

	if err := c.BindJSON(&candidate); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("UPDATE candidates SET name = ?, party_id = ? WHERE id = ?", candidate.Name, candidate.PartyID, id)
	if err != nil {
		log.Printf("Error updating data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	candidateWithID := gin.H{
		"id":       id,
		"name":     candidate.Name,
		"party_id": candidate.PartyID,
	}
	log.Printf("Candidate Updated: %v", candidateWithID)
	c.JSON(http.StatusOK, candidateWithID)
}

func DeleteCandidate(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM candidates WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	log.Printf("Candidate Deleted: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Candidate deleted"})
}
