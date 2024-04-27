package model

import "github.com/google/uuid"

type User struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	ProfilePic string    `json:"profile_pic"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
