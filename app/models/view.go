package models

type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Email             string `json:"email"`
	Name              string `json:"name"`
	EmployeeId        int64  `json:"employee_id"`
	Role              string `json:"role"`
	TokenId           string `json:"tokenId"`
	IsProfileComplete bool   `json:"isProfileComplete"`
}
