// controllers/article_controller.go
package controllers

import (
	"context"
	"net/http"
	"time"

	"Server/db"
	"Server/models"
	"Server/utils"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// Set estimated read time
	article.EstimatedReadTime = utils.EstimateReadingTime(article.Content)

	// Update the UID in Firestore
	_, err = docRef.Set(context.Background(), article)
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

func AddCommentOnArticle(c *gin.Context) {
	articleID := c.Param("id")

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Generate a unique comment ID
	comment.CommentID = uuid.New().String()
	comment.Commented = time.Now()

	// Update the article document to add the comment to the comments array
	_, err := db.FirestoreClient.Collection(articlesCollection).Doc(articleID).Update(context.Background(), []firestore.Update{
		{Path: "Comments", Value: firestore.ArrayUnion(comment)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}

func RemoveCommentOnArticle(c *gin.Context) {
	articleID := c.Param("id")
	commentID := c.Query("commentID")

	// Get the article document from Firestore
	articleRef := db.FirestoreClient.Collection("articles").Doc(articleID)
	articleSnap, err := articleRef.Get(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve article"})
		return
	}

	// Parse the article data into a struct
	var article models.Article
	if err := articleSnap.DataTo(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse article data"})
		return
	}

	// Find the index of the comment with the specified commentID
	var commentIndex = -1
	for i, comment := range article.Comments {
		if comment.CommentID == commentID {
			commentIndex = i
			break
		}
	}

	// Check if the commentID exists in the article's comments array
	if commentIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Remove the comment from the comments array
	article.Comments = append(article.Comments[:commentIndex], article.Comments[commentIndex+1:]...)

	// Update the article document back in Firestore
	_, err = articleRef.Set(context.Background(), article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment removed successfully"})
}
