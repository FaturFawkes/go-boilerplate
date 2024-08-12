package common

type BaseResponseApi[T any] struct {
	Success bool              `json:"success,omitempty"`
	Message string            `json:"message,omitempty"`
	Data    T                 `json:"data,omitempty"`
	Meta    *BaseMetaResponse `json:"meta,omitempty"`
}

type BaseMetaResponse struct {
	Count int `json:"count,omitempty"`
}
