package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Здесь хотим хранить функции работающие с обсуждениями
*/

type Discussion struct {
	ID           int64  `json:"id"`
	Production   int64  `json:"production"`
	Author       string `json:"author"`
	Topic        string `json:"topic"`
	EntryMessage string `json:"entry_message"`
	Message      string `json:"message"`
}

// компонента LastDiscussion. Хотим вернуть последний комментарий главного обсуждения.
func LastDiscussion(c *gin.Context, db *sql.DB) {
	query := `
        SELECT c.text, u.username, u.pfp
        FROM comments c
        JOIN users u ON u.username = c.author
        WHERE c.type = 'discussion' AND c.entity_id = '1'
        ORDER BY c.id DESC
        LIMIT 1
    `

	row := db.QueryRow(query)

	var text, username, pfp string

	err := row.Scan(&text, &username, &pfp)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": "no comments"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"text":     text,
		"username": username,
		"pfp":      pfp,
	})
}

// компонента MainPage. Возращаем id случайного обсуждения
func RandomDiscussion(c *gin.Context, db *sql.DB) {
	query := `
        SELECT id 
        FROM discussions 
        ORDER BY RANDOM() 
        LIMIT 1
    `

	row := db.QueryRow(query)

	var discussionID string

	err := row.Scan(&discussionID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": discussionID,
	})
}

// аналогична админке
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

// аналогична админке
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

// аналогична админке
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
