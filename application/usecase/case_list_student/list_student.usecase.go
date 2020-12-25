package case_list_student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-student-microservice/application/repository"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/helper"
)

//ListStudentUseCase is a contract
type ListStudentUseCase interface {
	ListStudent(c *gin.Context)
}

type listStudentUseCase struct {
	studentRepository repository.StudentRepository
}

//NewListStudentUseCase creates a new instance of ListStudentUseCase
func NewListStudentUseCase(repo repository.StudentRepository) ListStudentUseCase {
	return &listStudentUseCase{
		studentRepository: repo,
	}
}

func (ctl *listStudentUseCase) ListStudent(c *gin.Context) {
	students := ctl.studentRepository.ListStudent()
	helper.BuildResponse(http.StatusOK, students, c)
}
