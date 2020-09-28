package listing

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mauriciobergallo/school-students-api/pkg/logging"
)

var ErrInRepository = errors.New("There is an error in the Repository")
var ErrNotFound = errors.New("Student not found")

type Repository interface {
	GetStudents() ([]Student, error)
	GetStudentById(uuid.UUID) (Student, error)
}

// Service interface used to list the strings
type Service interface {
	ListStudents() ([]Student, error)
	ObtainStudentById(uuid.UUID) (Student, error)
}

type service struct{
	r Repository
	l logging.Service
}

// NewService constructor of the default service.
func NewService(r Repository, l logging.Service) Service {
	return &service{r, l}
}

// ListStudents returns the list of the Student registered.
func (s *service) ListStudents() ([]Student, error) {
	sl, err := s.r.GetStudents()
	if err != nil {
		s.l.Error(err.Error())
		return nil, ErrInRepository
	}

	return sl, nil
}

// ObtainStudentById returns an Student filtered by ID if the Student doesn't exist returns null.
func (s *service) ObtainStudentById(id uuid.UUID) (Student, error) {
	st, err := s.r.GetStudentById(id)
	if err != nil {
		s.l.Error(err.Error())
		return Student{}, ErrInRepository
	}

	if st.LastName == "" {
		return Student{}, ErrNotFound
	}

	return st, nil
}
