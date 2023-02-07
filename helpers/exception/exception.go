package exception

import "net/http"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
func Throw(err error, status int) {
	panic(&Error{
		Code:    status,
		Message: err.Error(),
	})
}

func NotFoundException(err ...error) {
	if len(err) > 0 {
		panic(&Error{
			Code:    http.StatusNotFound,
			Message: err[0].Error(),
		})

	}
}

func BadRequestException(err ...error) {
	if len(err) > 0 {
		panic(&Error{
			Code:    http.StatusBadRequest,
			Message: err[0].Error(),
		})

	}
}

func InternalServerErrorException(err ...error) {
	if len(err) > 0 {
		panic(&Error{
			Code:    http.StatusInternalServerError,
			Message: err[0].Error(),
		})
	}
}
