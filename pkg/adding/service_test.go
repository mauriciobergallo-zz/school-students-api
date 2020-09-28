package adding_test

import (
  "github.com/mauriciobergallo/school-students-api/pkg/adding"
  "github.com/mauriciobergallo/school-students-api/pkg/logging"
  "github.com/mauriciobergallo/school-students-api/pkg/storage/memory"
  "github.com/stretchr/testify/assert"
  "testing"
  "time"
)

type testService struct {
  r adding.Repository
  l logging.Service
}

var ts *testService

func setUp() {
  sm := new(memory.Storage)
  sm.RunSeeds()
  l := logging.NewStdoutLogging("DEBUG")

  ts = &testService{r: sm, l:l }
}

func tearDown() {
  //ts = nil
}

func TestMain(m *testing.M) {
  setUp()
  m.Run()
  tearDown()
}

func TestService_AddStudent_WithCorrectData_ShouldBeOk(t *testing.T) {
  // Given
  s := adding.NewService(ts.r, ts.l)
  st := adding.Student{
    FirstName:            "John",
    LastName:             "Doe",
    IdentificationNumber: "12345678",
    EMail:                "john.doe@mail.com",
    BirthDate:            time.Date(1980, 2, 15, 0, 0, 0, 0, time.UTC),
  }

  // When
  obtained, err := s.AddStudent(st)

  // Then
  assert.Nil(t, err, "Not error expected.")
  assert.NotNil(t, obtained, "The Student was not created.")
}
