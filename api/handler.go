package api

import (
	"context"
	"net/http"
	"time"

	"github.com/asinha24/verify-otp/data"
	"github.com/gin-gonic/gin"
)

const appTimeout = 10 * time.Second

func (app *Config) SendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, "OTP sent successfully")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyOTP
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.VerifyOTP{
			User: payload.User,
			Code: payload.Code,
		}

		if err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code); err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")
	}
}
