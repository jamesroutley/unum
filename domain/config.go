package domain

type Config struct {
	TwilioAccountSID string `json:"twilio_account_sid"`
	TwilioAuthToken  string `json:"twilio_auth_token"`
	PhoneNumber      string `json:"phone_number"`
}
