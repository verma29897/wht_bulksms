package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/verma29897/bulksms/models"
)

var AUTH_TOKEN = os.Getenv("META_ACCESS_TOKEN")
// Prefer using META_ACCESS_TOKEN everywhere for consistency

type Job struct {
	Request models.MessageRequest
	Contact string
}

func SendMessagesHandler(c *gin.Context) {
	var req models.MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := processWithWorkers(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send messages: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Messages sent successfully"})
}

func processWithWorkers(req models.MessageRequest) error {
	const workerCount = 10
	jobs := make(chan Job, len(req.ContactList))
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Feed jobs
	for _, contact := range req.ContactList {
		jobs <- Job{Request: req, Contact: contact}
	}
	close(jobs)

	wg.Wait()
	return nil
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		err := sendSingleMessage(job.Request, job.Contact)
		if err != nil {
			fmt.Printf("[Worker %d] Failed to send to %s: %v\n", id, job.Contact, err)
		} else {
			fmt.Printf("[Worker %d] Sent to %s successfully\n", id, job.Contact)
		}
	}
}

func sendSingleMessage(req models.MessageRequest, contact string) error {
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/messages", req.PhoneNumberID)

	headerComponent := map[string]interface{}{
		"type":       "header",
		"parameters": []interface{}{},
	}

	if req.MediaID != nil && (req.MediaType == "IMAGE" || req.MediaType == "DOCUMENT" || req.MediaType == "VIDEO" || req.MediaType == "AUDIO") {
		headerComponent["parameters"] = append(headerComponent["parameters"].([]interface{}), map[string]interface{}{
			"type": req.MediaType,
			req.MediaType: map[string]string{
				"id": *req.MediaID,
			},
		})
	}

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                contact,
		"type":              "template",
		"template": map[string]interface{}{
			"name":     req.TemplateName,
			"language": map[string]string{"code": req.Language},
			"components": []interface{}{
				headerComponent,
				map[string]interface{}{"type": "body", "parameters": []interface{}{}},
			},
		},
	}

	jsonData, _ := json.Marshal(payload)
	reqBody := bytes.NewBuffer(jsonData)

	httpReq, _ := http.NewRequest("POST", url, reqBody)
	httpReq.Header.Set("Authorization", "Bearer "+AUTH_TOKEN)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Meta API error: %s", string(bodyBytes))
	}
	return nil
}
