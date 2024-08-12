package response

type (
	Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}
)

func DefaultResponse(success bool, message string, data any) Response {
	return Response{
		Success: success,
		Message: message,
		Data: data,
	}
}
