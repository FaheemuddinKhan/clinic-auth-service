package dto

import "time"

//represents new user registration request
type NewUserRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	DOB       time.Time `json:"dob"`
	Street    string    `json:"street"`
	PinCode   string    `json:"pin_code"`
	City      string    `json:"city"`
	District  string    `json:"district"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
}

//represents new user registration response
type NewUserResponse struct {
	UserId string `json:"user_id"`
	Token string `json:"jwt_token"`
}