// controllers/blog_controller.go
package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"Server/db"
	"Server/models"
	"Server/utils"
)

const blogsCollection = "blogs"

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.Date = time.Now()

	docRef, _, err := db.FirestoreClient.Collection(blogsCollection).Add(context.Background(), blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	blog.ID = docRef.ID

	blog.EstimatedReadTime = utils.EstimateReadingTime(blog.Content)

	if _, err := docRef.Set(context.Background(), blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog with ID"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog created successfully", "id": blog.ID})
}

func GetBlog(c *gin.Context) {
	blogID := c.Param("id")

	doc, err := db.FirestoreClient.Collection(blogsCollection).Doc(blogID).Get(context.Background())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get blog"})
		return
	}

	var blog models.Blog
	if err := doc.DataTo(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse blog data"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func GetAllBlogs(c *gin.Context) {
	var blogs []models.Blog

	iter := db.FirestoreClient.Collection(blogsCollection).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blogs"})
			return
		}

		var blog models.Blog
		if err := doc.DataTo(&blog); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse blog data"})
			return
		}

		blogs = append(blogs, blog)
	}

	c.JSON(http.StatusOK, blogs)
}

func UpdateBlog(c *gin.Context) {
	blogID := c.Param("id")

	var updatedBlog models.Blog
	if err := c.ShouldBindJSON(&updatedBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBlog.ID = blogID

	_, err := db.FirestoreClient.Collection(blogsCollection).Doc(blogID).Set(context.Background(), updatedBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func DeleteBlog(c *gin.Context) {
	blogID := c.Param("id")

	_, err := db.FirestoreClient.Collection(blogsCollection).Doc(blogID).Delete(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
