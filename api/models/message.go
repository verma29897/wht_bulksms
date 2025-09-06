package models

type MessageRequest struct {
	PhoneNumberID string   `json:"phone_number_id"`
	TemplateName  string   `json:"template_name"`
	Language      string   `json:"language"`
	MediaType     string   `json:"media_type"`
	MediaID       *string  `json:"media_id"`
	ContactList   []string `json:"contact_list"`
}
