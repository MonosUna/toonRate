package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCollection(c *gin.Context, db *sql.DB) {
	var col Collection
	if err := c.ShouldBindJSON(&col); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `INSERT INTO collections (id, author, topic, message) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, col.ID, col.Author, col.Topic, col.Message)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func GetCollections(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, author, topic, message FROM collections")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var collections []Collection
	for rows.Next() {
		var col Collection
		if err := rows.Scan(&col.ID, &col.Author, &col.Topic, &col.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		collections = append(collections, col)
	}
	c.JSON(http.StatusOK, gin.H{"collections": collections})
}

func GetCollection(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var col Collection
	query := `SELECT id, author, topic, message FROM collections WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&col.ID, &col.Author, &col.Topic, &col.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}
	c.JSON(http.StatusOK, col)
}

func EditCollection(c *gin.Context, db *sql.DB) {
	var col Collection
	if err := c.ShouldBindJSON(&col); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `UPDATE collections SET author = $1, topic = $2, message = $3 WHERE id = $4`
	_, err := db.Exec(query, col.Author, col.Topic, col.Message, col.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func DeleteCollection(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	query := `DELETE FROM collections WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}
