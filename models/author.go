package models

import "time"

type Author struct {
	Id         string     `json:"id"`
	Firstname  string     `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John"`
	Lastname   string     `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe"`
	Middlename string     `json:"middlename" example:"Junior"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}

type CreateAuthorModel struct {
	Firstname  string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John"`
	Lastname   string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe"`
	Middlename string `json:"middlename" example:"Junior"`
}

type UpdateAuthorModel struct {
	Id        string `json:"id" binding:"required"`
	Firstname string `json:"firstname" minLength:"2" maxLength:"50"`
	Lastname  string `json:"lastname" minLength:"2" maxLength:"50"`
	Middlename string `json:"middlename"`
}
