package httpx

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type SuccessResponse[T any] struct {
	Data T `json:"data"`
}
