package dtos

import (
	"server/pkg/constant"

	"github.com/invopop/validation"
)

type (
	UpdateUserStatus struct {
		Status string `json:"status"`
	}

	UpdateUserStatusRequest struct {
		UserId int64
		UpdateUserStatus
	}
)

func (uus UpdateUserStatus) Validate() error {
	return validation.ValidateStruct(&uus,
		validation.Field(&uus.Status, validation.In(constant.UserActive, constant.UserInactive)),
	)
}
