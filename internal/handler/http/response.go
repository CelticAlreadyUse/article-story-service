package http_handler

type Response struct{
	Data any `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}
