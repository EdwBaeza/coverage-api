package infraestructure

type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Path       string `json:"path"`
}
