// handlers/upload.go
package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/verma29897/bulksms/utils"
)

func UploadMedia(c *gin.Context) {
	phoneNumberID := c.PostForm("phone_number_id")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer openedFile.Close()

	// Detect content type from extension
	ext := filepath.Ext(file.Filename)
	if len(ext) > 0 {
		ext = ext[1:] // remove dot
	}
	contentType := utils.GetMediaFormat(ext)

	// Prepare multipart form body
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	_ = writer.WriteField("type", contentType)
	_ = writer.WriteField("messaging_product", "whatsapp")

	part, _ := writer.CreateFormFile("file", file.Filename)
	_, _ = io.Copy(part, openedFile)
	writer.Close()

	// Prepare HTTP request
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/media", phoneNumberID)
	req, _ := http.NewRequest("POST", url, &b)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("META_ACCESS_TOKEN"))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Meta API upload failed"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}
