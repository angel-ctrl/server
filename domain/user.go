package domain

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Users struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Secret    string    `json:"secret"`
	UserState string    `json:"userState"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Plan      time.Time `json:"plan"`
}

func (u Users) Validate() error {
	return validation.ValidateStruct(&u,
		//validation.Field(&u.ID, validation.Required),
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.UserState, validation.Required),
		validation.Field(&u.Secret, validation.Required),
	)
}


func (u Users) Validate_without_pass() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.UserState, validation.Required),
		validation.Field(&u.Secret, validation.Required),
	)
}

func (u Users) Validate_without_pass_userstate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Secret, validation.Required),
	)
}
