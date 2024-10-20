package controllers

import (
	"net/http"
	"strconv"

	"github.com/asahi-zip/api/models"

	"github.com/gin-gonic/gin"
)

func UploadMedia(c *gin.Context) {
	orgIDStr := c.PostForm("org_id")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get file"})
		return
	}

	orgID, err := strconv.ParseUint(orgIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid org_id"})
		return
	}

	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer fileData.Close()

	fileContent := make([]byte, file.Size)
	fileData.Read(fileContent)

	media := models.Media{
		FileName: file.Filename,
		FileData: fileContent,
		FileSize: file.Size,
		OrgID:    uint(orgID),
	}

	if err := models.DB.Create(&media).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file in database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file uploaded successfully"})
}
