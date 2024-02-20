// controllers/video_controller.go
package controllers

import (
	"Server/db"
	"Server/models"
	"context"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

const videosCollection = "videos"

func CreateVideo(c *gin.Context) {
	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video.UploadDate = time.Now()

	docRef, _, err := db.FirestoreClient.Collection(videosCollection).Add(context.Background(), video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create video"})
		return
	}

	// Retrieve the generated UID from the Firestore document reference
	videoID := docRef.ID

	// Set the video ID in the model
	video.ID = videoID

	// Update the UID in Firestore
	_, err = docRef.Update(context.Background(), []firestore.Update{
		{Path: "ID", Value: videoID},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ID in Firestore"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Video created successfully", "id": videoID})
}

func GetVideo(c *gin.Context) {
	videoID := c.Param("id")

	doc, err := db.FirestoreClient.Collection(videosCollection).Doc(videoID).Get(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	var video models.Video
	doc.DataTo(&video)

	c.JSON(http.StatusOK, video)
}

func GetVideos(c *gin.Context) {
	var videos []models.Video

	iter := db.FirestoreClient.Collection(videosCollection).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
			return
		}

		var video models.Video
		doc.DataTo(&video)
		videos = append(videos, video)
	}

	c.JSON(http.StatusOK, videos)
}

func UpdateVideo(c *gin.Context) {
	videoID := c.Param("id")

	var updatedFields models.Video
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedFields.ID = videoID

	// Update the video with the updated fields
	_, err := db.FirestoreClient.Collection(videosCollection).Doc(videoID).Set(context.Background(), updatedFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video updated successfully"})
}

func DeleteVideo(c *gin.Context) {
	videoID := c.Param("id")

	_, err := db.FirestoreClient.Collection(videosCollection).Doc(videoID).Delete(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video deleted successfully"})
}

func AddCommentOnVideo(c *gin.Context) {
	videoID := c.Param("id")

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Generate a unique comment ID
	comment.CommentID = uuid.New().String()
	comment.Commented = time.Now()

	// Update the video document to add the comment to the comments array
	_, err := db.FirestoreClient.Collection(videosCollection).Doc(videoID).Update(context.Background(), []firestore.Update{
		{Path: "Comments", Value: firestore.ArrayUnion(comment)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}

func RemoveCommentOnVideo(c *gin.Context) {
	videoID := c.Param("id")
	commentID := c.Query("commentID")

	// Get the video document from Firestore
	videoRef := db.FirestoreClient.Collection(videosCollection).Doc(videoID)
	videoSnap, err := videoRef.Get(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video"})
		return
	}

	// Parse the video data into a struct
	var video models.Video
	if err := videoSnap.DataTo(&video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse video data"})
		return
	}

	// Find the index of the comment with the specified commentID
	var commentIndex = -1
	for i, comment := range video.Comments {
		if comment.CommentID == commentID {
			commentIndex = i
			break
		}
	}

	// Check if the commentID exists in the video's comments array
	if commentIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Remove the comment from the comments array
	video.Comments = append(video.Comments[:commentIndex], video.Comments[commentIndex+1:]...)

	// Update the video document back in Firestore
	_, err = videoRef.Set(context.Background(), video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment removed successfully"})
}

func LikeVideo(c *gin.Context) {
	videoID := c.Param("id")

	// Update the likes count of the video
	_, err := db.FirestoreClient.Collection("videos").Doc(videoID).Update(context.Background(), []firestore.Update{
		{Path: "Likes", Value: firestore.Increment(1)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video liked successfully"})
}
