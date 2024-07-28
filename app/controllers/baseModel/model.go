package baseModel

type ErrorResponse struct {
	IsError      bool   `json:"is_error"`
	ErrorMessage string `json:"error_message,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
}
