package repository

import (
	"fmt"

	"github.com/ydhnwb/elib-student-microservice/domain/entity"
	"gorm.io/gorm"
)

//StudentRepository is a contract
type StudentRepository interface {
	CreateStudent(s entity.Student) (entity.Student, error)
	UpdateStudent(s entity.Student) (entity.Student, error)
	ListStudent() []entity.Student
	SearchStudent(query string) []entity.Student
	DeleteStudent(studentID string) (bool, error)
}

type studentRepository struct {
	db *gorm.DB
}

//NewStudentRepository creates a new instance of StudentRepository
func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

func (repo *studentRepository) CreateStudent(s entity.Student) (entity.Student, error) {
	student := entity.Student{}
	err := repo.db.Save(&student).Error
	if err != nil {
		return student, err
	}
	return student, nil
}

func (repo *studentRepository) UpdateStudent(s entity.Student) (entity.Student, error) {
	err := repo.db.Save(&s).Error
	if err != nil {
		return s, err
	}
	return s, nil
}

func (repo *studentRepository) ListStudent() []entity.Student {
	students := []entity.Student{}
	repo.db.Find(&students)
	return students
}

func (repo *studentRepository) SearchStudent(query string) []entity.Student {
	students := []entity.Student{}
	repo.db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", query)).Find(&students)
	return students
}

func (repo *studentRepository) DeleteStudent(studentID string) (bool, error) {
	student := entity.Student{}
	err := repo.db.Where("id = ?", studentID).Take(&student).Error
	if err != nil {
		return false, err
	}
	repo.db.Delete(&student)
	return true, nil
}
