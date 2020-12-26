package case_detail_student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-student-microservice/application/repository"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/helper"
)

//DetailStudentUseCase is a contract
type DetailStudentUseCase interface {
	DetailStudent(c *gin.Context)
}

type detailStudentUseCase struct {
	studentRepository repository.StudentRepository
}

//NewDetailStudentUseCase creates a new instance of detail student
func NewDetailStudentUseCase(repo repository.StudentRepository) DetailStudentUseCase {
	return &detailStudentUseCase{
		studentRepository: repo,
	}
}

func (ctl detailStudentUseCase) DetailStudent(c *gin.Context) {
	id := c.Param("id")
	res := ctl.studentRepository.FindByID(id)
	helper.BuildResponse(http.StatusOK, res, c)
}
