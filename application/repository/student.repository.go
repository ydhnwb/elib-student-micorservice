package repository

import (
	"github.com/ydhnwb/elib-student-microservice/domain/entity"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/dto"
)

//StudentRepository is a contract
type StudentRepository interface {
	CreateStudent(u dto.StudentCreateDTO) (entity.Student, error)
	UpdateStudent(u dto.StudentUpdateDTO) (entity.Student, error)
	ListStudent() []entity.Student
	SearchStudent(query string) []entity.Student
	DeleteStudent(studentID string) interface{}
}
