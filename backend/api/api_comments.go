package api

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
	Здесь хотим хранить функции работающие с комментариями
*/

type Comment struct {
	ID       int64  `json:"id"`
	Type     string `json:"type"`
	Text     string `json:"text"`
	EntityID int64  `json:"entity_id"`
	Author   string `json:"author"`
}

// компонента Discussion/Review/Collection. Получаем все комментарии к произведению
func GetComments(c *gin.Context, db *sql.DB, commentType string) {
	idS := c.Param("id")
	id, _ := strconv.ParseInt(idS, 10, 64)

	query := `
		SELECT id, type, text, author
		FROM comments
		WHERE type = $1 AND entity_id = $2;
	`

	rows, err := db.Query(query, commentType, id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment
		comment.EntityID = id
		if err := rows.Scan(&comment.ID, &comment.Type, &comment.Text, &comment.Author); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// компонента Discussion/Review/Collection. Получаем все оценки комментария.
func GetCommentRating(c *gin.Context, db *sql.DB) {
	idS := c.Param("id")
	id, _ := strconv.ParseInt(idS, 10, 64)
	type Rating struct {
		ID     int64  `json:"id"`
		User   string `json:"user"`
		Rating int    `json:"rating"`
	}

	var ratings []Rating

	query := `
        SELECT id, author, rating
        FROM ratings
        WHERE entity_id = $1 AND type = 'comment'
    `
	rows, err := db.Query(query, id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rating Rating
		if err := rows.Scan(&rating.ID, &rating.User, &rating.Rating); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		ratings = append(ratings, rating)
	}

	if err := rows.Err(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ratings": ratings})
}

// компонента Discussion/Review/Collection. Получаем количество комментариев.
func GetCommentCount(c *gin.Context, db *sql.DB) {

	var count int
	query := `
        SELECT COUNT(*) 
        FROM comments 
    `
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment_count": count})
}

// компонента Discussion/Review/Collection. Добавляем новый комментарий.
func AddNewComment(c *gin.Context, db *sql.DB, commentType string) {
	var newComment Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	if newComment.Type == "" {
		newComment.Type = commentType
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	defer tx.Rollback()

	var maxID sql.NullInt64
	err = tx.QueryRow("SELECT MAX(id) FROM comments").Scan(&maxID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	var newID int
	if !maxID.Valid {
		newID = 1
	} else {
		newID = int(maxID.Int64) + 1
	}

	_, err = tx.Exec(
		"INSERT INTO comments (id, type, text, entity_id, author) VALUES ($1, $2, $3, $4, $5)",
		newID, newComment.Type, newComment.Text, newComment.EntityID, newComment.Author,
	)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// компонента Discussion/Review/Collection. лайк/дизлайк комментарию
func RateComment(c *gin.Context, db *sql.DB) {
	var newRating NewRatingOrStatus

	if err := c.ShouldBindJSON(&newRating); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	if newRating.Rating != -1 && newRating.Rating != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	defer tx.Rollback()

	var existingRatingID string
	query := `
		SELECT id FROM ratings
		WHERE entity_id = $1 AND author = $2 AND type = 'comment'
	`
	err = tx.QueryRow(query, newRating.ID, newRating.Author).Scan(&existingRatingID)

	if err == sql.ErrNoRows {
		var maxID sql.NullInt64
		err := tx.QueryRow("SELECT MAX(id) FROM ratings").Scan(&maxID)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		var newID int
		if !maxID.Valid {
			newID = 1
		} else {
			newID = int(maxID.Int64) + 1
		}

		insertQuery := `
			INSERT INTO ratings (id, type, rating, entity_id, author)
			VALUES ($1, 'comment', $2, $3, $4)
		`
		_, err = tx.Exec(insertQuery, newID, newRating.Rating, newRating.ID, newRating.Author)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		if err := tx.Commit(); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": ""})
		return
	} else if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	updateQuery := `
		UPDATE ratings 
		SET rating = $1 
		WHERE id = $2 
	`
	_, err = tx.Exec(updateQuery, newRating.Rating, existingRatingID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}
