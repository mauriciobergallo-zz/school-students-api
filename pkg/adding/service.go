package adding

import "github.com/mauriciobergallo/school-app-students-api/pkg/logging"

type Repository interface {
	InsertStudent(Student) (Student, error)
}

type Service interface {
	AddStudent(Student) (Student, error)
}

type service struct {
	r	Repository
	l logging.Service
}

func NewService(r Repository, l logging.Service) Service {
	return &service{r, l}
}

func (s *service) AddStudent(sn Student) (Student, error) {
	st, err := s.r.InsertStudent(sn)
	if err != nil {

	}

	return st, nil
}
