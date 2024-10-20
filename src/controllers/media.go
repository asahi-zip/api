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

	mimeType := file.Header.Get("Content-Type")

	media := models.Media{
		FileName: file.Filename,
		MimeType: mimeType,
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

func GetMedia(c *gin.Context) {
	mediaIDStr := c.Param("id")

	mediaID, err := strconv.ParseUint(mediaIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid media_id"})
		println(err)
		return
	}

	var media models.Media
	if err := models.DB.Where("id = ?", mediaID).First(&media).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	if media.MimeType == "" {
		media.MimeType = "application/octet-stream"
	}

	c.Data(http.StatusOK, media.MimeType, media.FileData)
}
