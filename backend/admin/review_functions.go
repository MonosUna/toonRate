package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddReview(c *gin.Context, db *sql.DB) {
	var r Review
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `INSERT INTO reviews (id, production_id, author, topic, message) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, r.ID, r.Production, r.Author, r.Topic, r.Message)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func GetReviews(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, production_id, author, topic, message FROM reviews")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var r Review
		if err := rows.Scan(&r.ID, &r.Production, &r.Author, &r.Topic, &r.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		reviews = append(reviews, r)
	}
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func GetReview(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var r Review
	query := `SELECT id, production_id, author, topic, message FROM reviews WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&r.ID, &r.Production, &r.Author, &r.Topic, &r.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}
	c.JSON(http.StatusOK, r)
}

func EditReview(c *gin.Context, db *sql.DB) {
	var r Review
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `UPDATE reviews SET production_id = $1, author = $2, topic = $3, message = $4 WHERE id = $5`
	_, err := db.Exec(query, r.Production, r.Author, r.Topic, r.Message, r.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func DeleteReview(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	query := `DELETE FROM reviews WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}
