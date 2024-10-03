package dtos

import (
	"github.com/invopop/validation"
)

type (
	CreateUserRequest struct {
		FullName    string `json:"fullName"`
		PhoneNumber string `json:"phoneNumber"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}

	CreateUserResponse struct {
		UserId    int64  `json:"userId"`
		Token     string `json:"token"`
		ExpiredAt int64  `json:"expiredAt"`
	}
)

func (cup CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&cup,
		validation.Field(&cup.FullName, validation.Required, validation.Length(0, 50)),
		validation.Field(&cup.PhoneNumber, validation.Required, validation.Length(0, 13)),
		validation.Field(&cup.Email, validation.Required),
		validation.Field(&cup.Password, validation.Required),
	)
}
