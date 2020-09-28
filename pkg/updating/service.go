package updating

import "github.com/mauriciobergallo/school-students-api/pkg/logging"

type Repository interface {
  UpdateStudent(Student) (Student, error)
}

type Service interface {
  ReplaceStudent(Student) (Student, error)
}

type service struct {
  r Repository
  l logging.Service
}

func NewService(r Repository, l logging.Service) Service {
  return &service{r,l}
}

func (s *service) ReplaceStudent(st Student) (Student, error) {
  ust, err := s.r.UpdateStudent(st)
  if err != nil {
    return st, err
  }

  return ust, nil
}
