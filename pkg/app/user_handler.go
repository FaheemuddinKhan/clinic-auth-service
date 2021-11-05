package app

import (
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/logger"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/service"
	"net/http"
)

type UserHandler struct {
	service service.IUserService
}

func (h *UserHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) CheckAvailable(rw http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	logger.Info(v.Encode())
	if len(v) != 1 {
		writeResponse(rw, http.StatusBadRequest, "Need exactly one query param {username | email | phone}")
	} else {
		username, ok := v["username"]
		if ok {
			available, err := h.service.CheckAvailable("username", username)
			if err != nil {
				writeResponse(rw, err.Code, err.Message)
			}
			writeResponse(rw, http.StatusOK, available)
			return
		}
		email, ok := v["email"]
		if ok {
			available, err := h.service.CheckAvailable("email", email)
			if err != nil {
				writeResponse(rw, err.Code, err.Message)
			}
			writeResponse(rw, http.StatusOK, available)
			return
		}
		phone, ok := v["email"]
		if ok {
			available, err := h.service.CheckAvailable("phone", phone)
			if err != nil {
				writeResponse(rw, err.Code, err.Message)
			}
			writeResponse(rw, http.StatusOK, available)
			return
		}
	}
}
