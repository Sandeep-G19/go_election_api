package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResults(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, election_id, candidate_id, votes FROM results")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	var results []gin.H
	for rows.Next() {
		var id, electionID, candidateID, votes int
		if err := rows.Scan(&id, &electionID, &candidateID, &votes); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		results = append(results, gin.H{
			"id":           id,
			"election_id":  electionID,
			"candidate_id": candidateID,
			"votes":        votes,
		})
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with rows"})
		return
	}

	log.Printf("Results: %v", results)
	c.JSON(http.StatusOK, results)
}

func GetResult(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	row := db.QueryRow("SELECT id, election_id, candidate_id, votes FROM results WHERE id = ?", id)

	var result gin.H
	var electionID, candidateID, votes int
	if err := row.Scan(&id, &electionID, &candidateID, &votes); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No result found with id: %v", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "No result found"})
		} else {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
		}
		return
	}
	result = gin.H{
		"id":           id,
		"election_id":  electionID,
		"candidate_id": candidateID,
		"votes":        votes,
	}
	log.Printf("Result: %v", result)
	c.JSON(http.StatusOK, result)
}

func CreateResult(c *gin.Context, db *sql.DB) {
	var result struct {
		ElectionID  int `json:"election_id"`
		CandidateID int `json:"candidate_id"`
		Votes       int `json:"votes"`
	}

	if err := c.BindJSON(&result); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	res, err := db.Exec("INSERT INTO results (election_id, candidate_id, votes) VALUES (?, ?, ?)", result.ElectionID, result.CandidateID, result.Votes)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	id, _ := res.LastInsertId()
	resultWithID := gin.H{
		"id":           id,
		"election_id":  result.ElectionID,
		"candidate_id": result.CandidateID,
		"votes":        result.Votes,
	}
	log.Printf("Result Created: %v", resultWithID)
	c.JSON(http.StatusCreated, resultWithID)
}

func UpdateResult(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var result struct {
		ElectionID  int `json:"election_id"`
		CandidateID int `json:"candidate_id"`
		Votes       int `json:"votes"`
	}

	if err := c.BindJSON(&result); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("UPDATE results SET election_id = ?, candidate_id = ?, votes = ? WHERE id = ?", result.ElectionID, result.CandidateID, result.Votes, id)
	if err != nil {
		log.Printf("Error updating data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	resultWithID := gin.H{
		"id":           id,
		"election_id":  result.ElectionID,
		"candidate_id": result.CandidateID,
		"votes":        result.Votes,
	}
	log.Printf("Result Updated: %v", resultWithID)
	c.JSON(http.StatusOK, resultWithID)
}

func DeleteResult(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM results WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	log.Printf("Result Deleted: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Result deleted"})
}
