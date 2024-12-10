package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GetEnvVariable(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func (app *Config) validateJsonBody(c *gin.Context, data any) error {
	var validate = validator.New()

	//validate the request body
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	//validate with the validator library
	if err := validate.Struct(&data); err != nil {
		return err
	}
	return nil
}

func (app *Config) writeJSON(c *gin.Context, status int, data any) {
	c.JSON(status, jsonResponse{Status: status, Message: "success", Data: data})
}

func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})
}
