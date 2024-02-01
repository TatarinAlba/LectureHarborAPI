package entity

import (
	"time"
)

type User struct {
	Id                    int64     `json:"user_id"`
	Username              []byte    `json:"username"`
	Email                 string    `json:"email"`
	HashedPassword        []byte    `json:"hashed_password"`
	RegistrationTimestamp time.Time `json:"registration_timestamp"`
}

type UserRequest struct {
	Username []byte `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password []byte `json:"password" validate:"required"`
}

type UserResponse struct {
	Message   string `json:"message_status"`
	JwtString string `json:"jwt_string"`
}
