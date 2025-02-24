package models

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResError[T any](message string, data T) Response[T] {
	return Response[T]{
		Success: false,
		Message: message,
		Data:    data,
	}
}

func ResSuccess[T any](data T) Response[T] {
	return Response[T]{
		Success: true,
		Message: "Success",
		Data:    data,
	}
}
