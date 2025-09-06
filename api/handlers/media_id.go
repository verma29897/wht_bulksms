package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	API_VERSION  = "v20.0"
    APP_ID       = ""
)

func UploadHeaderHandle(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// Save uploaded file to temp path
	tempPath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}
	defer os.Remove(tempPath)

	// Open the file
	openedFile, err := os.Open(tempPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer openedFile.Close()

	// Get file size
	stat, _ := openedFile.Stat()
	fileLength := stat.Size()
	ext := strings.ToLower(filepath.Ext(file.Filename)[1:])
	fileType := getMediaFormat(ext)

	// Step 1: Initiate Upload Session
    appID := os.Getenv("APP_ID")
    accessToken := os.Getenv("META_ACCESS_TOKEN")

    initURL := fmt.Sprintf("https://graph.facebook.com/%s/%s/uploads", API_VERSION, appID)
	req, _ := http.NewRequest("POST", initURL, nil)

	q := req.URL.Query()
	q.Add("file_length", fmt.Sprintf("%d", fileLength))
	q.Add("file_type", fileType)
    q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload session"})
		return
	}
	defer resp.Body.Close()

	var jsonResp map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode upload session response"})
		return
	}

	uploadID := fmt.Sprintf("%v", jsonResp["id"])

	// Step 2: Upload File to Upload URL
    uploadURL := fmt.Sprintf("https://graph.facebook.com/%s/%s", API_VERSION, uploadID)
	fileContent, _ := os.ReadFile(tempPath)

	uploadReq, _ := http.NewRequest("POST", uploadURL, bytes.NewReader(fileContent))
    uploadReq.Header.Set("Authorization", "OAuth "+accessToken)
	uploadReq.Header.Set("file_offset", "0")

	uploadResp, err := http.DefaultClient.Do(uploadReq)
	if err != nil || uploadResp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "upload failed"})
		return
	}
	defer uploadResp.Body.Close()

	var uploadResult map[string]interface{}
	if err := json.NewDecoder(uploadResp.Body).Decode(&uploadResult); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode upload result"})
		return
	}

	if hVal, ok := uploadResult["h"]; ok {
		c.JSON(http.StatusOK, gin.H{"header_handle": hVal})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing header handle in response"})
	}
}

// Media format detection
func getMediaFormat(ext string) string {
	formats := map[string]string{
		"jpg": "image/jpeg", "jpeg": "image/jpeg", "png": "image/png",
		"gif": "image/gif", "bmp": "image/bmp", "svg": "image/svg+xml",
		"mp4": "video/mp4", "avi": "video/x-msvideo", "mov": "video/quicktime",
		"flv": "video/x-flv", "mkv": "video/x-matroska", "mp3": "audio/mpeg",
		"aac": "audio/aac", "ogg": "audio/ogg", "wav": "audio/wav",
		"pdf": "application/pdf", "doc": "application/msword",
		"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"ppt":  "application/vnd.ms-powerpoint",
		"pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		"xls":  "application/vnd.ms-excel",
		"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"txt":  "text/plain", "csv": "text/csv",
	}
	if val, exists := formats[ext]; exists {
		return val
	}
	return "application/octet-stream"
}
