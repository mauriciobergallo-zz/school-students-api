package listing

import (
  "github.com/google/uuid"
  "time"
)

type Student struct {
  ID                   uuid.UUID `json:"id"`
  FirstName            string    `json:"firstName"`
  LastName             string    `json:"lastName"`
  IdentificationNumber string    `json:"identificationNumber"`
  EMail                string    `json:"email"`
  BirthDate            time.Time `json:"birthDate"`
}

func NewStudent(firstName string, lastName string, identificationNumber string, EMail string, birthDate time.Time) *Student {
  id, _ := uuid.NewRandom()
  return &Student{ID: id, FirstName: firstName, LastName: lastName, IdentificationNumber: identificationNumber, EMail: EMail, BirthDate: birthDate}
}
