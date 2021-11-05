package docs

import "github.com/faheemuddinkhan/clinic-auth-service/api/dto"

// swagger:route POST /phone/otp?type=request phoneotp phoneotprequest
// Requests for the otp to verify Phone number of user
// responses:
//   200: phoneotprequestresponse

//Phone OTP Request's response contains OTP id which has to be sent back along with the OTP received on phone
//swagger:response phoneotprequestresponse
type PhoneOTPRequestResponseWrapper struct{
	//in:body
	Body dto.NewPhoneOTPResponse
}

//swagger:parameters phoneotprequest
type PhoneOTPRequestParamsWrapper struct{
	//in:body
	Body dto.NewPhoneOTPRequest
	//in:query
	Type string `json:"type"`
}


// swagger:route POST /phone/otp?type=verify phoneotp phoneotpverify
// Verifies the otp of the user sent along this request
// responses:
//   200: phoneotpverifyresponse


//Response for the Verification of Phone OTP
//swagger:response phoneotpverifyresponse
type PhoneOTPVerifyResponseWrapper struct{
	//in:body
	Body dto.VerifyOTPResponse
}

//swagger:parameters phoneotpverify
type PhoneOTPVerifyParamsWrapper struct{
	//in:body
	Body dto.VerifyOTPRequest
	//in:query
	Type string `json:"type"`
}