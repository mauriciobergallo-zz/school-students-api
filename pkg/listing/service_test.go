package listing

import (
	"github.com/google/uuid"
	"github.com/mauriciobergallo/school-app-students-api/pkg/logging"
	"testing"
	"time"
)

func setUp() {
	ls = logging.NewStdoutLogging("DEBUG")
	mr = new(mockRepository)

	idM, _ := uuid.Parse("a7f8feed-383a-4213-989e-268411297104")
	ns1 := Student{
		ID:                   idM,
		FirstName:            "Mauricio",
		LastName:             "Bergallo",
		IdentificationNumber: "31174646",
		EMail:                "bergallo.mauricio@gmail.com",
		BirthDate:            time.Date(1984, 9, 12, 0, 0, 0, 0, time.UTC),
	}

	mr.students = append(mr.students, ns1)

	idC, _ := uuid.Parse("81b81256-2cac-419b-8fec-d2d69b0a0dd2")
	ns2 := Student{
		ID:                   idC,
		FirstName:            "Cintia",
		LastName:             "D'Allotta",
		IdentificationNumber: "308762323",
		EMail:                "cinty.dallotta@gmail.com",
		BirthDate:            time.Date(1984, 3, 4, 0, 0, 0, 0, time.UTC),
	}

	mr.students = append(mr.students, ns2)
}

func tearDown() {

}

func TestMain(m *testing.M) {
	setUp()
	m.Run()
	tearDown()
}

type mockRepository struct {
	students []Student
}

var mr *mockRepository
var ls logging.Service

func TestService_GetStudents(t *testing.T) {
	// Given
	s := NewService(mr, ls)
	expectedIdMauricio, _ := uuid.Parse("a7f8feed-383a-4213-989e-268411297104")
	expectedIdCintia, _ := uuid.Parse("81b81256-2cac-419b-8fec-d2d69b0a0dd2")

	// When
	obtained, err := s.ListStudents()

	// Then
	if err != nil {
		t.Errorf("Not error expected, but got: %q", err.Error())
	}

	if len(obtained) != 2 {
		t.Errorf("Did not get expected result. Wanted: 2, got: #{len(obtained)}")
	}

	if obtained[0].ID != expectedIdMauricio {
		t.Errorf("Did not get expected result. Wanted: %q, got: %q", expectedIdMauricio, obtained[0].ID)
	}

	if obtained[1].ID != expectedIdCintia {
		t.Errorf("Did not get expected result. Wanted: %q, got: %q", expectedIdCintia, obtained[0].ID)
	}
}

func (mr *mockRepository) GetStudents() ([]Student, error) {
	return mr.students, nil
}

func (mr *mockRepository) GetStudentById(id uuid.UUID) (Student, error) {
	s := Student{}

	for i := range mr.students {
		if mr.students[i].ID == id {
			return s, nil
		}
	}

	return s, nil
}
