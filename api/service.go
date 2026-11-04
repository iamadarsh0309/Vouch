package api 

import(
	"errors"
	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest/verify/v2"
)

var twilioClient *twilio.RestClient

func init(){
	twilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: envACCOUNTSID(),
		Password: envAUTHTOKEN(),
	})
}

func(app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	
res, err := twilioClient.Verify.Services(envSERVICESID()).Verifications.Create(params)
	if err != nil {
		return "", err
	}
	return res.Sid, nil
}

func(app *Config) twilioVerifyOTP(phoneNumber string, code string) (string, error) {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)
	res, err := twilioClient.Verify.Services(envSERVICESID()).Verifications.Create(params)
	if err != nil {
		return "", err
	}
	return res.Sid, nil
}