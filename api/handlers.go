package api

import (
	"context"
	"go-appwriteSDK/data"
	"net/http"
	"time"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/gin-gonic/gin"
)

var projectId = GetEnvVariable("PROJECT_ID")
var databaseId = GetEnvVariable("DATABASE_ID")
var collectionId = GetEnvVariable("COLLECTION_ID")
var apiKey = GetEnvVariable("API_KEY")
var endpoint = GetEnvVariable("API_URL")

var client = appwrite.NewClient(
	appwrite.WithEndpoint(endpoint),
	appwrite.WithProject(projectId),
	appwrite.WithKey(apiKey),
)

const appTimeout = time.Second * 10

func (app *Config) createProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		databases := appwrite.NewDatabases(client)
		var payload data.ProjectRequest
		defer cancel()

		app.validateJsonBody(ctx, &payload)

		doc, err := databases.CreateDocument(databaseId, collectionId, id.Unique(), payload)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusCreated, doc)
	}
}

func (app *Config) getProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		databases := appwrite.NewDatabases(client)
		docId := ctx.Param("projectId")
		defer cancel()

		response, err := databases.GetDocument(databaseId, collectionId, docId)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		var document data.Project
		err = response.Decode(&document)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusOK, document)
	}
}

func (app *Config) updateProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		databases := appwrite.NewDatabases(client)
		var payload data.ProjectRequest
		docId := ctx.Param("projectId")
		defer cancel()

		app.validateJsonBody(ctx, &payload)
		updates := data.ProjectRequest{
			Name:        payload.Name,
			Description: payload.Description,
		}

		response, err := databases.UpdateDocument(databaseId, collectionId, docId, databases.WithUpdateDocumentData(updates))
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		var document data.Project
		err = response.Decode(&document)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusOK, document)
	}
}

func (app *Config) deleteProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		databases := appwrite.NewDatabases(client)
		docId := ctx.Param("projectId")
		defer cancel()

		response, err := databases.DeleteDocument(databaseId, collectionId, docId)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusOK, response)
	}
}
