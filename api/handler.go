package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RunRequest struct {
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

type RunOutput struct {
	Output string `json:"output"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func RunCodeHandler(c *gin.Context) {
	var request RunRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request body: " + err.Error(),
		})

		return
	}

	if !isValidLanguage(request.Language) {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Unsupported programming language: " + request.Language,
		})

		return
	}

	userId, _ := c.Get("userId")

	output, err := runCode(request, userId.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, RunOutput{
		Output: output,
	})
}
