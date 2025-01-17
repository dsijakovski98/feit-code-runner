package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dsijakovski98/feit-code-runner/languages"
	"github.com/dsijakovski98/feit-code-runner/utils"
	"github.com/gin-gonic/gin"
)

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

func CleanupDebugHandler(c *gin.Context) {
	var request CleanupRequest

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

	if !supportsTests(request.Language) {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Programming language" + request.Language + " does not support tests!",
		})

		return
	}

	langExtension := languages.ProgrammingLanguages[request.Language].GetConfig().Extension

	filename := fmt.Sprintf("%s.%s", request.TaskName, langExtension)
	filePath, err := utils.CreateTestFile(filename, request.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
	}

	defer os.Remove(filePath)

	if err := cleanupCode(request, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	cleanCode := string(content)

	c.JSON(http.StatusOK, CleanupResponse{
		CleanCode: cleanCode,
	})
}
