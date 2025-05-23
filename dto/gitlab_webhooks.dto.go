package dto

import "time"

// GitlabWebhooks ..
type GitlabWebhooks struct {
	ObjectKind        string      `json:"object_kind"`
	EventName         string      `json:"event_name"`
	Before            string      `json:"before"`
	After             string      `json:"after"`
	Ref               string      `json:"ref"`
	CheckoutSha       string      `json:"checkout_sha"`
	Message           interface{} `json:"message"`
	UserID            int         `json:"user_id"`
	UserName          string      `json:"user_name"`
	UserUsername      string      `json:"user_username"`
	UserEmail         string      `json:"user_email"`
	UserAvatar        string      `json:"user_avatar"`
	ProjectID         int         `json:"project_id"`
	Project           Project     `json:"project"`
	Commits           []Commits   `json:"commits"`
	TotalCommitsCount int         `json:"total_commits_count"`
	PushOptions       PushOptions `json:"push_options"`
	Repository        Repository  `json:"repository"`
}
type Project struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	WebURL            string      `json:"web_url"`
	AvatarURL         interface{} `json:"avatar_url"`
	GitSSHURL         string      `json:"git_ssh_url"`
	GitHTTPURL        string      `json:"git_http_url"`
	Namespace         string      `json:"namespace"`
	VisibilityLevel   int         `json:"visibility_level"`
	PathWithNamespace string      `json:"path_with_namespace"`
	DefaultBranch     string      `json:"default_branch"`
	CiConfigPath      interface{} `json:"ci_config_path"`
	Homepage          string      `json:"homepage"`
	URL               string      `json:"url"`
	SSHURL            string      `json:"ssh_url"`
	HTTPURL           string      `json:"http_url"`
}
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Commits struct {
	ID        string        `json:"id"`
	Message   string        `json:"message"`
	Title     string        `json:"title"`
	Timestamp time.Time     `json:"timestamp"`
	URL       string        `json:"url"`
	Author    Author        `json:"author"`
	Added     []string      `json:"added"`
	Modified  []string      `json:"modified"`
	Removed   []interface{} `json:"removed"`
}
type PushOptions struct {
}
type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}
