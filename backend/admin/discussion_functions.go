package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddDiscussion(c *gin.Context, db *sql.DB) {
	var d Discussion
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	query := `INSERT INTO discussions (id, production_id, author, topic, entry_message, message)
				VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, d.ID, d.Production, d.Author, d.Topic, d.EntryMessage, d.Message)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func GetDiscussions(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, production_id, author, topic, entry_message, message FROM discussions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var discussions []Discussion
	for rows.Next() {
		var d Discussion
		if err := rows.Scan(&d.ID, &d.Production, &d.Author, &d.Topic, &d.EntryMessage, &d.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		discussions = append(discussions, d)
	}
	c.JSON(http.StatusOK, gin.H{"discussions": discussions})
}

func GetDiscussion(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var d Discussion
	query := `SELECT id, production_id, author, topic, entry_message, message FROM discussions WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&d.ID, &d.Production, &d.Author, &d.Topic, &d.EntryMessage, &d.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}
	c.JSON(http.StatusOK, d)
}

func EditDiscussion(c *gin.Context, db *sql.DB) {
	var d Discussion
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `UPDATE discussions SET production_id = $1, author = $2, topic = $3, entry_message = $4, message = $5 WHERE id = $6`
	_, err := db.Exec(query, d.Production, d.Author, d.Topic, d.EntryMessage, d.Message, d.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func DeleteDiscussion(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	query := `DELETE FROM discussions WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}
