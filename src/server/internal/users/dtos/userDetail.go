package dtos

import "time"

type (
	UserDetailResponse struct {
		UserID      int64     `json:"id"`
		Email       string    `json:"email"`
		Fullname    string    `json:"fullName"`
		PhoneNumber string    `json:"phoneNumber"`
		UserType    string    `json:"userType"`
		IsActive    bool      `json:"isActive"`
		CreatedAt   time.Time `json:"createdAt"`
	}
)
