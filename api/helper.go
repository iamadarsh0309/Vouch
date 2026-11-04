package api 

import(
	"github.com/gin-gonic/gin"
	validate "github.com/go-playground/validator/v10"
)

type JSONResponse struct {
	status int `json:"status"`
	message string `json:"message"`
	data any `json:"data,omitempty"`
}

var v = validate.New()

func(app *Config) validateBody(c *gin.Context, payload any) error {
	if err := c.ShouldBindJSON(payload); err != nil {
		return err
	}
	if err := v.Struct(payload); err != nil {
		return err
	}
	return nil
}

func(app *Config) writeJSON(c *gin.Context, status int, data any, headers http.Header) error {
	return c.JSON(status, data)
}

func(app *Config) errorJSON(c *gin.Context, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	return c.JSON(statusCode, JSONResponse{
		status: statusCode,
		message: err.Error(),
		data: nil,
	})
}