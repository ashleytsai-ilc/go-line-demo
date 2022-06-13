package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type PushMessage struct {
	UserID  string `json:"userId" binding:"required,min=20,max=80"`
	Content string `json:"content" binding:"required,min=1,trim"`
}

type GetMessages struct {
	UserID string `uri:"userid" binding:"required,min=20,max=80"`
}

var Trim validator.Func = func(fl validator.FieldLevel) bool {
	return strings.Trim(fl.Field().String(), " ") != ""
}
