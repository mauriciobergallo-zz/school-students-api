package deleting

import "github.com/mauriciobergallo/school-app-students-api/pkg/logging"

type Repository interface {
  DeleteStudent(Student) error
}

type Service interface {
  RemoveStudent(Student) error
}

type service struct {
  r Repository
  l logging.Service
}

func NewService(r Repository, l logging.Service) Service {
 return &service{r, l}
}

func (s *service) RemoveStudent(st Student) error {
  err := s.r.DeleteStudent(st)
  if err != nil {
    return err
  }

  return nil
}
