package memory

import (
  "github.com/google/uuid"
  "time"
)

type Student struct {
  ID                   uuid.UUID
  FirstName            string
  LastName             string
  IdentificationNumber string
  EMail                string
  BirthDate            time.Time
  Created              time.Time
}


