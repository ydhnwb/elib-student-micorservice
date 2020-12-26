package case_create_student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/ydhnwb/elib-student-microservice/application/repository"
	"github.com/ydhnwb/elib-student-microservice/domain/entity"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/dto"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/helper"
)

//CreateStudentUseCase act as handler/controller
//It is more spesific than regular service
type CreateStudentUseCase interface {
	CreateStudent(c *gin.Context)
}

type createStudentUseCase struct {
	studentRepository repository.StudentRepository
}

//NewCreateStudentUseCase creates a new instance of
func NewCreateStudentUseCase(repo repository.StudentRepository) CreateStudentUseCase {
	return &createStudentUseCase{
		studentRepository: repo,
	}
}

func (ctl *createStudentUseCase) CreateStudent(c *gin.Context) {
	studentCreateDTO := dto.StudentCreateDTO{}
	e := c.ShouldBind(&studentCreateDTO)
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, c)
		return
	}

	student := entity.Student{}
	e = smapping.FillStruct(&student, smapping.MapFields(&studentCreateDTO))
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, c)
		return
	}

	res, err := ctl.studentRepository.CreateStudent(student)
	if err != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, c)
		return
	}

	helper.BuildResponse(http.StatusCreated, res, c)
}
