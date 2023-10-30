package error_utils

type CustomError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) GetStatus() int {
	return e.Status
}

func Cek() error {
	return &CustomError{
		Message: "Error",
		Status:  500,
	}
}
