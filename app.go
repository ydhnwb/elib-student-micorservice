package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-student-microservice/application/repository"
	"github.com/ydhnwb/elib-student-microservice/application/usecase/case_create_student"
	"github.com/ydhnwb/elib-student-microservice/application/usecase/case_delete_student"
	"github.com/ydhnwb/elib-student-microservice/application/usecase/case_detail_student"
	"github.com/ydhnwb/elib-student-microservice/application/usecase/case_list_student"
	"github.com/ydhnwb/elib-student-microservice/application/usecase/case_search_student"
	"github.com/ydhnwb/elib-student-microservice/application/usecase/case_update_student"
	"github.com/ydhnwb/elib-student-microservice/infrastructure/persistence"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                                 = persistence.SetupDatabaseConnection()
	studentRepository repository.StudentRepository             = repository.NewStudentRepository(db)
	createStudent     case_create_student.CreateStudentUseCase = case_create_student.NewCreateStudentUseCase(studentRepository)
	updateStudent     case_update_student.UpdateStudentUseCase = case_update_student.NewUpdateStudentUseCase(studentRepository)
	deleteStudent     case_delete_student.DeleteStudentUseCase = case_delete_student.NewDeleteStudentUseCase(studentRepository)
	listStudent       case_list_student.ListStudentUseCase     = case_list_student.NewListStudentUseCase(studentRepository)
	searchStudent     case_search_student.SearchStudentUseCase = case_search_student.NewSearchStudentUseCase(studentRepository)
	detailStudent     case_detail_student.DetailStudentUseCase = case_detail_student.NewDetailStudentUseCase(studentRepository)
)

func main() {
	defer persistence.CloseDatabaseConnection(db)
	r := gin.Default()

	r.GET("api/search", searchStudent.Search)

	studentRoutes := r.Group("api/students")
	{
		studentRoutes.GET("/", listStudent.ListStudent)
		studentRoutes.POST("/", createStudent.CreateStudent)
		studentRoutes.GET("/:id", detailStudent.DetailStudent)
		studentRoutes.PUT("/:id", updateStudent.UpdateStudent)
		studentRoutes.DELETE("/:id", deleteStudent.DeleteStudent)

	}

	r.Run(":8081")
}
