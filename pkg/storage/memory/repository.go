package memory

import (
  "errors"
  "github.com/google/uuid"
  "github.com/mauriciobergallo/school-students-api/pkg/adding"
  "github.com/mauriciobergallo/school-students-api/pkg/deleting"
  "github.com/mauriciobergallo/school-students-api/pkg/listing"
  "github.com/mauriciobergallo/school-students-api/pkg/updating"
  "time"
)

type Storage struct {
  students []Student
}

func (m *Storage) RunSeeds() error {
  idM, _ := uuid.Parse("a7f8feed-383a-4213-989e-268411297104")
  ns1 := Student{
    ID:                   idM,
    FirstName:            "John",
    LastName:             "Doe",
    IdentificationNumber: "12345678",
    EMail:                "john.dow@mail.com",
    BirthDate:            time.Date(1984, 5, 11, 0, 0, 0, 0, time.UTC),
    Created:              time.Now(),
  }

  m.students = append(m.students, ns1)

  idC, _ := uuid.Parse("81b81256-2cac-419b-8fec-d2d69b0a0dd2")
  ns2 := Student{
    ID:                   idC,
    FirstName:            "Jane",
    LastName:             "Dow",
    IdentificationNumber: "45678912",
    EMail:                "jane.doe@mail.com",
    BirthDate:            time.Date(1984, 1, 8, 0, 0, 0, 0, time.UTC),
    Created:              time.Now(),
  }

  m.students = append(m.students, ns2)

  return nil
}

func (m *Storage) GetStudents() ([]listing.Student, error) {
  var sl []listing.Student

  for i := range m.students {
    s := listing.Student{
      ID:                   m.students[i].ID,
      FirstName:            m.students[i].FirstName,
      LastName:             m.students[i].LastName,
      IdentificationNumber: m.students[i].IdentificationNumber,
      EMail:                m.students[i].EMail,
      BirthDate:            m.students[i].BirthDate,
    }

    sl = append(sl, s)
  }

  return sl, nil
}

func (m *Storage) GetStudentById(id uuid.UUID) (listing.Student, error) {
  s := listing.Student{}
  for i := range m.students {
    if m.students[i].ID == id {
      s := listing.Student{
        ID:                   m.students[i].ID,
        FirstName:            m.students[i].FirstName,
        LastName:             m.students[i].LastName,
        IdentificationNumber: m.students[i].IdentificationNumber,
        EMail:                m.students[i].EMail,
        BirthDate:            m.students[i].BirthDate,
      }

      return s, nil
    }
  }

  return s, nil
}

func (m *Storage) InsertStudent(ns adding.Student) (adding.Student, error) {
  newStudent := Student{
    ID:                   uuid.New(),
    FirstName:            ns.FirstName,
    LastName:             ns.LastName,
    IdentificationNumber: ns.IdentificationNumber,
    EMail:                ns.EMail,
    BirthDate:            ns.BirthDate,
    Created:              time.Now(),
  }

  m.students = append(m.students, newStudent)

  return ns, nil
}

func (m *Storage) DeleteStudent(st deleting.Student) error {
  for i := range m.students {
    if m.students[i].ID == st.ID {
      m.students = append(m.students[:i], m.students[i+1:]...)

      return nil
    }
  }

  return errors.New("Student not found")
}

func (m *Storage) UpdateStudent(st updating.Student) (updating.Student, error) {
  for i := range m.students {
    if m.students[i].ID == st.ID {
      m.students[i].FirstName = st.FirstName
      m.students[i].LastName = st.LastName
      m.students[i].BirthDate = st.BirthDate
      m.students[i].EMail = st.EMail
      m.students[i].IdentificationNumber = st.IdentificationNumber

      return st, nil
    }
  }

  return st, errors.New("Student not found")
}