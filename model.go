package main

import "time"

type Person struct {
	Firstname string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John"`
	Lastname  string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe"`
}

type Content struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type Article struct {
	Id         string     `json:"id"`
	Content    Content    `json:"content"`
	Author     Person     `json:"author"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}

type CreateArticleModel struct {
	Content    Content    `json:"content"`
	Author     Person     `json:"author"`
}

type UpdateArticleModel struct {
	Id         string     `json:"id"`
	Content    Content    `json:"content"`
	Author     Person     `json:"author"`
}

type JSONResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONError struct {
	Error string `json:"error"`
}
