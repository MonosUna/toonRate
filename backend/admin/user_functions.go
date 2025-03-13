package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context, db *sql.DB) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `INSERT INTO users (id, username, password, email, pfp, description) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, u.ID, u.Username, u.Password, u.Email, u.Pfp, u.Description)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func GetUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, username, password, email, pfp, description FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Pfp, &u.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		users = append(users, u)
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var u User
	query := `SELECT id, username, email, password, pfp, description FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Pfp, &u.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}
	c.JSON(http.StatusOK, u)
}

func EditUser(c *gin.Context, db *sql.DB) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `UPDATE users SET username = $1, email = $2, password = $3, pfp = $4, description = $5 WHERE id = $6`
	_, err := db.Exec(query, u.Username, u.Password, u.Email, u.Pfp, u.Description, u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func DeleteUser(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}
