package response

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ExceptionError(c *fiber.Ctx, err error, status_code int) error {
	if err != nil {
		return FailOnError(c, err, status_code)
	} else {
		return nil
	}
}

type ErrorResponse struct {
	Error            string        `json:"error"`
	ErrorStatus      int           `json:"errorStatus"`
	ErrorDescription string        `json:"errorDescription"`
	ErrorAt          time.Time     `json:"errorAt"`
	ErrorTraceId     string        `json:"errorTraceId"`
	ErrorUri         string        `json:"errorUri"`
	ErrorFields      []interface{} `json:"errorFields"`
	ErrorData        string        `json:"errorData"`
	State            interface{}   `json:"state"`
}

func FailOnError(c *fiber.Ctx, err error, status int, field ...[]interface{}) error {
	var errorFields []interface{}
	if len(field) > 0 {
		errorFields = field[0]
	} else {
		errorFields = nil
	}
	return c.Status(status).JSON(ErrorResponse{
		Error:            http.StatusText(status),
		ErrorStatus:      status,
		ErrorDescription: err.Error(),
		ErrorAt:          time.Now(),
		ErrorTraceId:     c.GetRespHeader("X-Request-Id"),
		ErrorUri:         "",
		ErrorFields:      errorFields,
		ErrorData:        "",
		State:            nil,
	})
}
