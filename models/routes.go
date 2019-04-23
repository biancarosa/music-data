package models

//APIResponse defines a default API Response
type APIResponse struct {
	Data    *interface{} `json:"data,omitempty"`
	Message *string      `json:"message,omitempty"`
}
