package deleting_test

import (
  "github.com/google/uuid"
  "github.com/mauriciobergallo/school-students-api/pkg/deleting"
  "github.com/mauriciobergallo/school-students-api/pkg/logging"
  "github.com/mauriciobergallo/school-students-api/pkg/storage/memory"
  "testing"
)

type testService struct {
  r deleting.Repository
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

func TestService_RemoveStudent_WithNonExistentStudent_ShouldFail(t *testing.T) {
  // Given
  s := deleting.NewService(ts.r, ts.l)
  id, _ := uuid.Parse("e8f814e0-01d5-4a67-bf58-952638ee5ed9")
  st := deleting.Student{ID: id}

  // When
  err := s.RemoveStudent(st)

  // Then
  if err == nil {
    t.Errorf("Error expected, but got none")
  }

  if err.Error() != "Student not found" {
    t.Errorf("Not the error expected, got: %q", err.Error())
  }
}

func TestService_RemoveStudent_WithExistentStudent_ShouldBeOk(t *testing.T) {
  // Given
  s := deleting.NewService(ts.r, ts.l)
  id, _ := uuid.Parse("a7f8feed-383a-4213-989e-268411297104")
  st := deleting.Student{ID: id}

  // When
  err := s.RemoveStudent(st)

  // Then
  if err != nil {
    t.Errorf("Not error expected, but got: %q", err.Error())
  }
}
