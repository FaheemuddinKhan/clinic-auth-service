package service

import (
	"github.com/faheemuddinkhan/clinic-auth-service/api/dto"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/errors"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/jwtoken"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/logger"
)


//DrivenPort: Interface for PhoneService
type IPhoneService interface {
	RequestOTP(dto.NewPhoneOTPRequest) (*dto.NewPhoneOTPResponse, *errors.AppError)
	VerifyOTP(dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, *errors.AppError)
}


//Adapter: Phone Service
type DefaultPhoneService struct {
}

func (s DefaultPhoneService) RequestOTP(request dto.NewPhoneOTPRequest) (*dto.NewPhoneOTPResponse, *errors.AppError) {
	logger.Info("OTP Requested" + request.Phone)
	return &dto.NewPhoneOTPResponse{OTPId: "1298198291829182"}, nil
}

func (s DefaultPhoneService) VerifyOTP(request dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, *errors.AppError) {
	logger.Info("OTP Verification Request" + request.OTP)
	if request.OTP == "123456" {
		tokenString, err := jwtoken.CreateRegToken(request.Phone)
		if err != nil {
			return nil, err
		}
		return &dto.VerifyOTPResponse{
			Success:         true,
			RegistrationJWT: *tokenString,
		}, nil
	}
	return &dto.VerifyOTPResponse{
		Success:         false,
		RegistrationJWT: "",
	}, nil
}

func NewPhoneService() DefaultPhoneService {
	return DefaultPhoneService{}
}

