package case_update_student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/ydhnwb/elib-student-microservice/application/repository"
	"github.com/ydhnwb/elib-student-microservice/domain/entity"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/dto"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/helper"
)

//UpdateStudentUseCase is a contract
type UpdateStudentUseCase interface {
	UpdateStudent(c *gin.Context)
}

type updateStudentUseCase struct {
	studentRepository repository.StudentRepository
}

//NewUpdateStudentUseCase creates a new instance...
func NewUpdateStudentUseCase(repo repository.StudentRepository) UpdateStudentUseCase {
	return &updateStudentUseCase{
		studentRepository: repo,
	}
}

func (ctl *updateStudentUseCase) UpdateStudent(c *gin.Context) {
	u := dto.StudentUpdateDTO{}
	e := c.ShouldBind(&u)
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, c)
		return
	}
	student := entity.Student{}
	err := smapping.FillStruct(&student, smapping.MapFields(&u))
	if err != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, err.Error(), helper.EmptyObj{}, c)
		return
	}

	res, er := ctl.studentRepository.UpdateStudent(student)
	if er != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, er.Error(), helper.EmptyObj{}, c)
		return
	}
	helper.BuildResponse(http.StatusOK, res, c)
}
