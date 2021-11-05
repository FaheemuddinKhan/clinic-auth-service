package jwtoken

import (
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/errors"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/logger"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type PhoneHolderClaim struct {
	Phone string `json:"username"`
	jwt.StandardClaims
}

var jwtRegKey []byte

func init() {
	jwtRegKey = []byte("v5L0ganWflKMMGdKE4NM")
}

func CreateRegToken(Phone string) (*string, *errors.AppError) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &PhoneHolderClaim{
		Phone: Phone,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtRegKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		logger.Error("jwt could not be created" + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected JWT error")
	}

	return &tokenString, nil
}

func VerifyRegToken(tokenString string) (bool, *errors.AppError) {
	claims := &PhoneHolderClaim{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtRegKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, nil
		}
		logger.Error("Error while verifying Registration JWT" + err.Error())
		return false, errors.NewUnexpectedError("Unexpected JWT error")
	}
	if !tkn.Valid {
		return false, errors.NewNotAuthorizedError("Not Authorised")
	}

	return true, nil
}
