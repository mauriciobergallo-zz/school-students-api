package updating

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
