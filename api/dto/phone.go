package dto

//represents body of New Phone OTP Request
type NewPhoneOTPRequest struct {
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
}

//represents body of New Phone otp Response
type NewPhoneOTPResponse struct {
	OTPId string `json:"otp_id"`
}

//represents body of Verify otp Request
type VerifyOTPRequest struct {
	Phone string `json:"phone"`
	OTPId string `json:"otp_id"`
	OTP   string `json:"otp_value"`
}

//represents body of verify otp response
type VerifyOTPResponse struct {
	Success         bool   `json:"success"`
	RegistrationJWT string `json:"registration_jwt,omitempty"`
}
