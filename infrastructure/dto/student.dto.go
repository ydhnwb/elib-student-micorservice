package dto

//StudentCreateDTO is a structure that clients need to fill when creating a new student
type StudentCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	NIM  string `json:"nim" form:"nim" binding:"required"`
}

//StudentUpdateDTO is a structure that clients need to fill when updading a new student
type StudentUpdateDTO struct {
	ID   uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	NIM  string `json:"nim" form:"nim" binding:"required"`
}
