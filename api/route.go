package api 

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes(){
	app.Router.POST("/api/send-otp", app.SendOTP)
	app.Router.POST("/api/verify-otp", app.VerifyOTP)
}