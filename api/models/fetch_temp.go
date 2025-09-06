package models

type Template struct {
	TemplateID       string      `json:"template_id"`
	TemplateLanguage string      `json:"template_language"`
	TemplateName     string      `json:"template_name"`
	MediaType        string      `json:"media_type"`
	MediaLink        string      `json:"media_link"`
	Status           string      `json:"status"`
	Category         string      `json:"category"`
	TemplateData     string      `json:"template_data"`
	Button           interface{} `json:"button,omitempty"`
}
