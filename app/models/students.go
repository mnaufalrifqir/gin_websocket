package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Nim     string
	Name    string
	Jurusan string
}

type StudentResponse struct {
	ID        uint   `json:"ID"`
	CreatedAt string `json:"Created_at"`
	UpdatedAt string `json:"Updated_at"`
	DeletedAt string `json:"Deleted_at"`
	Nim       string `json:"Nim"`
	Name      string `json:"Name"`
	Jurusan   string `json:"Jurusan"`
}

type AllStudentResponse struct {
	Students []StudentResponse `json:"students"`
}

type StudentResponseByID struct {
	Student StudentResponse `json:"student"`
}

type StudentInput struct {
	Nim     string `json:"nim"`
	Name    string `json:"name"`
	Jurusan string `json:"jurusan"`
}

type StudentUpdate struct {
	Name string `json:"name"`
}

type ResponseStudentNotFound struct {
	Message string `json:"message" example:"Student not found!"`
}

type ResponseInvalid struct {
	Message string `json:"message" example:"Invalid request!"`
}

type ResponseSuccessCreate struct {
	Message string `json:"message" example:"Successfully create student!"`
}

type ResponseSuccessDelete struct {
	Message string `json:"message" example:"Successfully delete student!"`
}

type ResponseSuccessUpdate struct {
	Message string `json:"message" example:"Successfully update student!"`
}

type ResponseFailedCreate struct {
	Message string `json:"message" example:"Failed create student!"`
}

type ResponseFailedDelete struct {
	Message string `json:"message" example:"Failed delete student!"`
}

type ResponseFailedUpdate struct {
	Message string `json:"message" example:"Failed update student!"`
}
