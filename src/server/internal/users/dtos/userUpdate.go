package dtos

import (
	"server/pkg/constant"

	"github.com/invopop/validation"
)

type (
	UpdateUser struct {
		Fullname    string `json:"fullName"`
		PhoneNumber string `json:"phoneNumber"`
		UserType    string `json:"userType"`
	}

	UpdateUserRequest struct {
		UserId int64
		UpdateUser
	}
)

func (cup UpdateUser) Validate() error {
	return validation.ValidateStruct(&cup,
		validation.Field(&cup.UserType, validation.In(constant.UserTypePremium, constant.UserTypeRegular)),
	)
}
