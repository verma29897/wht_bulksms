// handlers/template.go
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/verma29897/bulksms/models"
)

func CreateTemplate(c *gin.Context) {
	var req models.TemplateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Build components
	components := []map[string]interface{}{}

	// HEADER
	if req.HeaderType != "" && req.HeaderContent != "" {
		switch req.HeaderType {
		case "headerText":
			components = append(components, map[string]interface{}{
				"type":   "HEADER",
				"format": "TEXT",
				"text":   req.HeaderContent,
			})
		case "headerImage", "headerVideo", "headerDocument", "headerAudio":
			format := req.HeaderType[len("header"):]
			components = append(components, map[string]interface{}{
				"type":   "HEADER",
				"format": format,
				"example": map[string]interface{}{
					"header_handle": []string{req.HeaderContent},
				},
			})
		}
	}

	// BODY
	components = append(components, map[string]interface{}{
		"type": "BODY",
		"text": req.BodyText,
	})

	// FOOTER
	if req.FooterText != "" {
		components = append(components, map[string]interface{}{
			"type": "FOOTER",
			"text": req.FooterText,
		})
	}

	// BUTTONS
	buttons := []map[string]interface{}{}
	if req.CallButtonText != "" && req.PhoneNumber != "" {
		buttons = append(buttons, map[string]interface{}{
			"type":         "PHONE_NUMBER",
			"text":         req.CallButtonText,
			"phone_number": req.PhoneNumber,
		})
	}
	if req.URLButtonText != "" && req.WebsiteURL != "" {
		buttons = append(buttons, map[string]interface{}{
			"type": "URL",
			"text": req.URLButtonText,
			"url":  req.WebsiteURL,
		})
	}
	if len(buttons) > 0 {
		components = append(components, map[string]interface{}{
			"type":    "BUTTONS",
			"buttons": buttons,
		})
	}

	// Build final payload
	payload := map[string]interface{}{
		"name":       req.TemplateName,
		"language":   req.Language,
		"category":   req.Category,
		"components": components,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal payload"})
		return
	}

	// Meta Graph API request
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/message_templates", req.WabaID)
	reqMeta, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	reqMeta.Header.Set("Content-Type", "application/json")
	reqMeta.Header.Set("Authorization", "Bearer "+os.Getenv("META_ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(reqMeta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request to Meta API failed"})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}
