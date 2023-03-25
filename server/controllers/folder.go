package controllers

import (
	"context"
	"memorly/configs"
	"memorly/forms"

	"memorly/models"
	"memorly/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var folderCollection *mongo.Collection = configs.GetCollection(configs.DB, "folders")

func CreateFolder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var folder forms.CreateFolderForm

		defer cancel()
		if err := c.BindJSON(&folder); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "Binding Error", Data: nil})
			return
		}

		if validationErr := validate.Struct(&folder); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "Validation Error", Data: nil})
			return
		}

		var user models.User

		id, _ := c.Get("id")
		err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusNotFound, responses.Response{Status: http.StatusNotFound, Message: "Not valid User", Data: nil})
			return
		}

		count, err := folderCollection.CountDocuments(ctx, bson.M{"title": folder.Title, "creatorId": user.Id})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "Error", Data: nil})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "Folder Name Already Taken", Data: nil})
			return
		}

		newFolder := models.Folder{
			Id:        primitive.NewObjectID(),
			Title:     folder.Title,
			CreatorId: user.Id,
		}

		newFolder.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newFolder.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		_, err = folderCollection.InsertOne(ctx, newFolder)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "Database Error", Data: nil})
			return
		}

		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "Success", Data: map[string]interface{}{"id": newFolder.Id}})
	}
}
