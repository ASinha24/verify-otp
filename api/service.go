package api

import (
	"errors"

	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: viper.GetString("TWILIO_ACCOUNT_SID"),
	Password: viper.GetString("TWILIO_AUTHTOKEN"),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{
		To: &phoneNumber,
	}
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(viper.GetString("TWILIO_SERVICE_SID"), params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func (app *Config) twilioVerifyOTP(phoneNumber, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{
		To:   &phoneNumber,
		Code: &code,
	}

	resp, err := client.VerifyV2.CreateVerificationCheck(viper.GetString("TWILIO_SERVICE_SID"), params)
	if err != nil {
		return err
	}

	if *resp.Status != "approved" {
		return errors.New("not a valid code")
	}

	return nil
}
