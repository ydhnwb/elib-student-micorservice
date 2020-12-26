package case_search_student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-student-microservice/application/repository"
	"github.com/ydhnwb/elib-student-microservice/domain/entity"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/helper"
)

//SearchStudentUseCase is a contract...
type SearchStudentUseCase interface {
	Search(c *gin.Context)
}

type searchStudentUseCase struct {
	studentRepository repository.StudentRepository
}

//NewSearchStudentUseCase creates a new instance of SearchStudentUseCase
func NewSearchStudentUseCase(repo repository.StudentRepository) SearchStudentUseCase {
	return &searchStudentUseCase{
		studentRepository: repo,
	}
}

func (ctl *searchStudentUseCase) Search(c *gin.Context) {
	query := c.Query("q")
	students := []entity.Student{}
	if query != "" {
		students = ctl.studentRepository.SearchStudent(query)
	} else {
		students = ctl.studentRepository.ListStudent()
	}
	helper.BuildResponse(http.StatusOK, students, c)
}
