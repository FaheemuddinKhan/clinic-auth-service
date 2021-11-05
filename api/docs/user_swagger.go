package docs

import "github.com/faheemuddinkhan/clinic-auth-service/api/dto"

// swagger:route POST /user/ user createuserrequest
// Creates new user
// responses:
//   200: createuserresponse

//The response contains success=true on successful user creation
//swagger:response createuserresponse
type NewUserCreateResponseWrapper struct{
	//in:body
	Body dto.NewUserResponse
}

//swagger:parameters createuserrequest
type NewUserCreateRequestWrapper struct{
	//in:body
	Body dto.NewUserRequest
}

// swagger:route GET /user/available user isavailablerequest
// Checks if object (username, email or phone) is available
// responses:
//   200: isavailableresponse

//The response contains true or false on availability of the object, one query at a time will be served
//swagger:response isavailableresponse
type IsAvailableResponseWrapper struct{
	//in:body
	Body bool
}

//swagger:parameters isavailablerequest
type IsAvailableRequestWrapper struct{
	//in:query
	Username string `json:"username,omitempty"`
	Email string	`json:"email,omitempty"`
	Phone string	`json:"phone,omitempty"`
}



