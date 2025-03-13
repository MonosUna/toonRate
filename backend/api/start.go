package api

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Start() {
	connStr := "user=postgres password=password123 dbname=fullstack sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/top_5_productions", func(c *gin.Context) {
		Top5Productions(c, db)
	})
	r.GET("/api/top_productions", func(c *gin.Context) {
		TopProductions(c, db)
	})
	r.GET("/api/production/:id", func(c *gin.Context) {
		GetProduction(c, db)
	})
	r.GET("/api/get_productions", func(c *gin.Context) {
		GetProductions(c, db)
	})
	r.GET("/api/get_users", func(c *gin.Context) {
		GetUsers(c, db)
	})
	r.POST("/api/add_user", func(c *gin.Context) {
		AddUser(c, db)
	})
	r.GET("/api/get_production_ratings/:id", func(c *gin.Context) {
		GetProductionRatings(c, db)
	})
	r.POST("/api/set_new_production_rating", func(c *gin.Context) {
		SetNewProductionRating(c, db)
	})
	r.POST("/api/set_new_production_status", func(c *gin.Context) {
		SetNewProductionStatus(c, db)
	})
	r.GET("/api/get_production_status/:id/:username", func(c *gin.Context) {
		GetProductionStatus(c, db)
	})
	r.GET("/api/get_statistics/:username", func(c *gin.Context) {
		Statistics(c, db)
	})
	r.POST("/api/update_user_description", func(c *gin.Context) {
		UpdateUserDescription(c, db)
	})
	r.POST("/api/update_user_pfp", func(c *gin.Context) {
		UpdateUserPfp(c, db)
	})
	r.GET("/api/all_product_status/:username", func(c *gin.Context) {
		AllProductionStatus(c, db)
	})
	r.GET("/api/get_discussions", func(c *gin.Context) {
		GetDiscussions(c, db)
	})
	r.POST("/api/add_discussion", func(c *gin.Context) {
		AddDiscussion(c, db)
	})
	r.GET("/api/get_discussion/:id", func(c *gin.Context) {
		GetDiscussion(c, db)
	})
	r.GET("/api/get_discussion_comments/:id", func(c *gin.Context) {
		GetComments(c, db, "discussion")
	})
	r.POST("/api/add_discussion_comment", func(c *gin.Context) {
		AddNewComment(c, db, "discussion")
	})
	r.GET("/api/get_comment_rating/:id", func(c *gin.Context) {
		GetCommentRating(c, db)
	})
	r.POST("/api/update_comment_rating", func(c *gin.Context) {
		RateComment(c, db)
	})
	r.GET("/api/count_of_comments", func(c *gin.Context) {
		GetCommentCount(c, db)
	})
	r.POST("/api/add_review", func(c *gin.Context) {
		AddReview(c, db)
	})
	r.GET("/api/get_reviews", func(c *gin.Context) {
		GetReviews(c, db)
	})
	r.GET("/api/get_review/:id", func(c *gin.Context) {
		GetReview(c, db)
	})
	r.GET("/api/get_review_comments/:id", func(c *gin.Context) {
		GetComments(c, db, "review")
	})
	r.GET("/api/get_collections", func(c *gin.Context) {
		GetCollections(c, db)
	})
	r.GET("/api/get_collections_productions", func(c *gin.Context) {
		GetCollectionsProductions(c, db)
	})
	r.POST("/api/add_review_comment", func(c *gin.Context) {
		AddNewComment(c, db, "review")
	})
	r.POST("/api/add_collection", func(c *gin.Context) {
		AddCollection(c, db)
	})
	r.POST("/api/add_collection_production", func(c *gin.Context) {
		AddCollectionProduction(c, db)
	})
	r.GET("/api/get_collection_comments/:id", func(c *gin.Context) {
		GetComments(c, db, "collection")
	})
	r.GET("/api/get_collection/:id", func(c *gin.Context) {
		GetCollection(c, db)
	})
	r.GET("/api/get_collection_productions/:id", func(c *gin.Context) {
		GetCollectionProductions(c, db)
	})
	r.GET("/api/random_production", func(c *gin.Context) {
		RandomProduction(c, db)
	})
	r.GET("/api/random_discussion", func(c *gin.Context) {
		RandomDiscussion(c, db)
	})
	r.GET("/api/random_review", func(c *gin.Context) {
		RandomReview(c, db)
	})
	r.GET("/api/random_collection", func(c *gin.Context) {
		RandomCollection(c, db)
	})
	r.GET("/api/last_discussion", func(c *gin.Context) {
		LastDiscussion(c, db)
	})
	r.POST("/api/add_collection_comment", func(c *gin.Context) {
		AddNewComment(c, db, "collection")
	})

	r.Run(":5050")
}
