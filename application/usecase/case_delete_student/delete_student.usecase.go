package case_delete_student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-student-microservice/application/repository"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/helper"
)

//DeleteStudentUseCase is a contract
type DeleteStudentUseCase interface {
	DeleteStudent(c *gin.Context)
}

type deleteStudentUseCase struct {
	studentRepository repository.StudentRepository
}

//NewDeleteStudentUseCase creates a new handler of delete user
func NewDeleteStudentUseCase(repo repository.StudentRepository) DeleteStudentUseCase {
	return &deleteStudentUseCase{
		studentRepository: repo,
	}
}

func (ctl *deleteStudentUseCase) DeleteStudent(c *gin.Context) {
	studentID := c.Param("id")
	_, err := ctl.studentRepository.DeleteStudent(studentID)
	if err == nil {
		helper.BuildResponse(http.StatusOK, helper.EmptyObj{}, c)
	} else {
		helper.BuildErrorResponse(http.StatusBadRequest, err.Error(), helper.EmptyObj{}, c)
	}
}
