package api 

import(
	"context"
	"net/http"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	 "github.com/adarsh/Vouch/data"
)
func(app *Config) SendOTP() gin.HandlerFunc{
	return func(c *gin.Context){
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		msg, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		
		app.writeJSON(c, http.StatusOK, gin.H{"message": "OTP sent successfully"})
	}
}

func(app *Config) VerifyOTP() gin.HandlerFunc{
	return func(c *gin.Context){
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(c, &payload)

		app.twilioVerifyOTP(payload.User.PhoneNumber, payload.Code)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		app.writeJSON(c, http.StatusOK, gin.H{"message": "OTP verified successfully"})
	}
}