package dto

// WebhookRequest ..
type WebhookRequest struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds,omitempty"`
}

// Embed ..
type Embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url,omitempty"`
	Color       int    `json:"color,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Footer      struct {
		IconURL string `json:"icon_url,omitempty"`
		Text    string `json:"text,omitempty"`
	} `json:"footer,omitempty"`
	Thumbnail struct {
		URL string `json:"url,omitempty"`
	} `json:"thumbnail,omitempty"`
	Image struct {
		URL string `json:"url,omitempty"`
	} `json:"image,omitempty"`
	Author struct {
		Name    string `json:"name,omitempty"`
		URL     string `json:"url,omitempty"`
		IconURL string `json:"icon_url,omitempty"`
	} `json:"author,omitempty"`
	Fields []struct {
		Name   string `json:"name,omitempty"`
		Value  string `json:"value,omitempty"`
		Inline bool   `json:"inline,omitempty"`
	} `json:"fields,omitempty"`
}
