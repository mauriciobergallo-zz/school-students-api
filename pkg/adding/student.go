package adding

import "time"

type Student struct {
  FirstName            string    `json:"firstName"`
  LastName             string    `json:"lastName"`
  IdentificationNumber string    `json:"identificationNumber"`
  EMail                string    `json:"email"`
  BirthDate            time.Time `json:"birthDate"`
}
