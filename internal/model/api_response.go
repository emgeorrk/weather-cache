package model

type APIResponse struct {
	Response Weather `json:"response,omitzero"`
	APIError
}

type APIError struct {
	Code    int    `json:"code,omitzero"`
	Error   string `json:"error,omitzero"`
	Message string `json:"message,omitzero"`
}
