// models/template.go
package models

type TemplateRequest struct {
	WabaID         string `json:"waba_id"`
	TemplateName   string `json:"template_name"`
	Language       string `json:"language"`
	Category       string `json:"category"`
	HeaderType     string `json:"header_type"`
	HeaderContent  string `json:"header_content"`
	BodyText       string `json:"body_text"`
	FooterText     string `json:"footer_text"`
	CallButtonText string `json:"call_button_text"`
	PhoneNumber    string `json:"phone_number"`
	URLButtonText  string `json:"url_button_text"`
	WebsiteURL     string `json:"website_url"`
}
