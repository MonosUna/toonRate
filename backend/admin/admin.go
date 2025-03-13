package admin

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Launch() {
	connStr := "user=postgres password=password123 dbname=fullstack sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/admin/add-production", func(c *gin.Context) {
		AddProduction(c, db)
	})
	r.GET("/admin/get-productions", func(c *gin.Context) {
		GetProductions(c, db)
	})
	r.DELETE("/admin/delete-production/:id", func(c *gin.Context) {
		DeleteProduction(c, db)
	})
	r.GET("/admin/get-production/:id", func(c *gin.Context) {
		GetProduction(c, db)
	})
	r.POST("/admin/update-production", func(c *gin.Context) {
		EditProduction(c, db)
	})

	r.POST("/admin/add-user", func(c *gin.Context) {
		AddUser(c, db)
	})
	r.GET("/admin/get-users", func(c *gin.Context) {
		GetUsers(c, db)
	})
	r.GET("/admin/get-user/:id", func(c *gin.Context) {
		GetUser(c, db)
	})
	r.POST("/admin/update-user", func(c *gin.Context) {
		EditUser(c, db)
	})
	r.DELETE("/admin/delete-user/:id", func(c *gin.Context) {
		DeleteUser(c, db)
	})

	r.POST("/admin/add-collection", func(c *gin.Context) {
		AddCollection(c, db)
	})
	r.GET("/admin/get-collections", func(c *gin.Context) {
		GetCollections(c, db)
	})
	r.GET("/admin/get-collection/:id", func(c *gin.Context) {
		GetCollection(c, db)
	})
	r.POST("/admin/update-collection", func(c *gin.Context) {
		EditCollection(c, db)
	})
	r.DELETE("/admin/delete-collection/:id", func(c *gin.Context) {
		DeleteCollection(c, db)
	})

	r.POST("/admin/add-discussion", func(c *gin.Context) {
		AddDiscussion(c, db)
	})
	r.GET("/admin/get-discussions", func(c *gin.Context) {
		GetDiscussions(c, db)
	})
	r.GET("/admin/get-discussion/:id", func(c *gin.Context) {
		GetDiscussion(c, db)
	})
	r.POST("/admin/update-discussion", func(c *gin.Context) {
		EditDiscussion(c, db)
	})
	r.DELETE("/admin/delete-discussion/:id", func(c *gin.Context) {
		DeleteDiscussion(c, db)
	})

	r.POST("/admin/add-review", func(c *gin.Context) {
		AddReview(c, db)
	})
	r.GET("/admin/get-reviews", func(c *gin.Context) {
		GetReviews(c, db)
	})
	r.GET("/admin/get-review/:id", func(c *gin.Context) {
		GetReview(c, db)
	})
	r.POST("/admin/update-review", func(c *gin.Context) {
		EditReview(c, db)
	})
	r.DELETE("/admin/delete-review/:id", func(c *gin.Context) {
		DeleteReview(c, db)
	})

	r.Run(":5050")
}
