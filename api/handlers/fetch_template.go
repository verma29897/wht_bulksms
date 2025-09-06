// handlers/template.go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/verma29897/bulksms/models"
)

func FetchTemplates(c *gin.Context) {
	wabaID := c.Param("waba_id")
	if wabaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing waba_id"})
		return
	}

	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/message_templates", wabaID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	q := req.URL.Query()
	q.Add("access_token", os.Getenv("META_ACCESS_TOKEN"))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Meta API request failed"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		c.JSON(resp.StatusCode, gin.H{"error": string(body)})
		return
	}

	var apiResp struct {
		Data []map[string]interface{} `json:"data"`
	}
	if err := json.Unmarshal(body, &apiResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	var templates []models.Template
	for _, entry := range apiResp.Data {
		// Extract top-level fields
		tmpl := models.Template{
			TemplateID:       entry["id"].(string),
			TemplateLanguage: entry["language"].(string),
			TemplateName:     entry["name"].(string),
			Status:           entry["status"].(string),
			Category:         entry["category"].(string),
		}

		// Loop through components
		if comps, ok := entry["components"].([]interface{}); ok {
			for _, comp := range comps {
				cmp := comp.(map[string]interface{})
				switch cmp["type"] {
				case "BODY":
					if text, ok := cmp["text"].(string); ok {
						tmpl.TemplateData = text
					}
				case "HEADER":
					if format, ok := cmp["format"].(string); ok {
						tmpl.MediaType = format
					}
					if example, ok := cmp["example"].(map[string]interface{}); ok {
						if handle, ok := example["header_handle"].([]interface{}); ok && len(handle) > 0 {
							tmpl.MediaLink = handle[0].(string)
						}
					}
				case "BUTTONS":
					tmpl.Button = cmp["buttons"]
				}
			}
		}

		templates = append(templates, tmpl)
	}

	c.JSON(http.StatusOK, gin.H{"templates": templates})
}
