package app

import (
	"encoding/json"
	"github.com/faheemuddinkhan/clinic-auth-service/api/dto"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/logger"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/service"
	"net/http"
	"time"
)

type PhoneHandler struct {
	service service.IPhoneService
}

func (h *PhoneHandler) PhoneOtp(rw http.ResponseWriter, r *http.Request) {
	logger.Info("In PhoneOtp")
	v := r.URL.Query()

	if v.Get("type") == "request" {
		h.RequestOtp(rw, r)
	} else if v.Get("type") == "verify" {
		h.VerifyOtp(rw, r)
	} else {
		writeResponse(rw, http.StatusBadRequest, "Query Param/Value not valid")
	}
}

func (h *PhoneHandler) RequestOtp(rw http.ResponseWriter, r *http.Request) {
	logger.Info("In Request Otp")
	var requestData dto.NewPhoneOTPRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		responseData, appError := h.service.RequestOTP(requestData)
		if appError != nil {
			writeResponse(rw, appError.Code, appError.Message)
		} else {
			writeResponse(rw, http.StatusCreated, responseData)
		}
	}
}

func (h *PhoneHandler) VerifyOtp(rw http.ResponseWriter, r *http.Request) {
	logger.Info("In Verify Otp")

	var requestData dto.VerifyOTPRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		responseData, appError := h.service.VerifyOTP(requestData)
		if appError != nil {
			writeResponse(rw, appError.Code, appError.Message)
		} else {
			http.SetCookie(rw, &http.Cookie{
				Name:    "reg_token",
				Value:   responseData.RegistrationJWT,
				Expires: time.Now().Add(4 * time.Minute),
			})
			responseData.RegistrationJWT = ""
			writeResponse(rw, http.StatusCreated, responseData)
		}
	}
}
