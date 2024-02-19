// controllers/article_controller.go
package controllers

import (
	"context"
	"net/http"
	"time"

	"Server/db"
	"Server/models"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

const articlesCollection = "articles"

func CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.UploadDate = time.Now()

	// Add the article to Firestore
	docRef, _, err := db.FirestoreClient.Collection(articlesCollection).Add(context.Background(), article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	// Retrieve the generated UID from the Firestore document reference
	articleID := docRef.ID

	// Set the article ID in the model
	article.ID = articleID

	// Update the UID in Firestore
	_, err = docRef.Update(context.Background(), []firestore.Update{
		{Path: "ID", Value: articleID},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ID in Firestore"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article created successfully", "id": articleID})
}

func GetArticle(c *gin.Context) {
	articleID := c.Param("id")

	doc, err := db.FirestoreClient.Collection(articlesCollection).Doc(articleID).Get(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	var article models.Article
	doc.DataTo(&article)

	c.JSON(http.StatusOK, article)
}

func GetAllArticles(c *gin.Context) {
	var articles []models.Article

	iter := db.FirestoreClient.Collection(articlesCollection).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve articles"})
			return
		}

		var article models.Article
		doc.DataTo(&article)
		articles = append(articles, article)
	}

	c.JSON(http.StatusOK, articles)
}

func UpdateArticle(c *gin.Context) {
	articleID := c.Param("id")

	var updatedArticle models.Article
	if err := c.ShouldBindJSON(&updatedArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the updated article has a valid ID
	updatedArticle.ID = articleID

	// Update the article in Firestore
	_, err := db.FirestoreClient.Collection(articlesCollection).Doc(articleID).Set(context.Background(), updatedArticle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully"})
}

func DeleteArticle(c *gin.Context) {
	articleID := c.Param("id")

	_, err := db.FirestoreClient.Collection(articlesCollection).Doc(articleID).Delete(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

func AddComment(c *gin.Context) {
	articleID := c.Param("id")

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set commented time
	comment.Commented = time.Now()

	// Add comment to the article in Firestore
	_, err := db.FirestoreClient.Collection(articlesCollection).Doc(articleID).Update(context.Background(), []firestore.Update{
		{Path: "comments", Value: firestore.ArrayUnion(comment)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}

func RemoveComment(c *gin.Context) {
	articleID := c.Param("id")
	commentID := c.Param("commentID")

	// Remove comment from article in Firestore
	_, err := db.FirestoreClient.Collection(articlesCollection).Doc(articleID).Update(context.Background(), []firestore.Update{
		{Path: "comments", Value: firestore.ArrayRemove(commentID)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment removed successfully"})
}

func LikeArticle(c *gin.Context) {
	articleID := c.Param("id")

	// Get user ID (assuming it's available in the context)
	userID := "USER_ID_PLACEHOLDER" // Example: retrieving user ID from middleware

	// Update article's likedBy array with the user's ID
	_, err := db.FirestoreClient.Collection(articlesCollection).Doc(articleID).Update(context.Background(), []firestore.Update{
		{Path: "likedBy", Value: firestore.ArrayUnion(userID)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article liked successfully"})
}

func UnlikeArticle(c *gin.Context) {
	articleID := c.Param("id")

	// Get user ID (assuming it's available in the context)
	userID := "USER_ID_PLACEHOLDER" // Example: retrieving user ID from middleware

	// Update article's likedBy array by removing the user's ID
	_, err := db.FirestoreClient.Collection(articlesCollection).Doc(articleID).Update(context.Background(), []firestore.Update{
		{Path: "likedBy", Value: firestore.ArrayRemove(userID)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article unliked successfully"})
}
