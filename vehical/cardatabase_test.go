package vehical

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestSet(t *testing.T) {
	testcases := []struct {
		desc           string
		input          Car
		expectedOutput bool
	}{

		{"valid values", Car{1, "Safari", "MGN1", "Petrol"}, true},
		{"valid values", Car{2, "Wagner", "FHG", "Petrol"}, true},
		{"valid values", Car{3, "Benz", "MGN1", "Diesel"}, true},
		{"valid values", Car{3, "Benz", "C-Class", "Diesel"}, false},
		{"valid values", Car{4, "XUV", "MGN2", "Petrol"}, true},
	}
	var s Store
	var err error
	db, mock, err := sqlmock.New()
	s.db = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// now we execute our method
	for i, value := range testcases {
		if value.expectedOutput == true {
			mock.ExpectExec("insert into Car values").
				WithArgs(value.input.Id, value.input.Name, value.input.Model, value.input.EngineType).
				WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(err)
		} else {
			mock.ExpectExec("insert into Car values").
				WithArgs(value.input.Id, value.input.Name, value.input.Model, value.input.EngineType).
				WillReturnResult(sqlmock.NewResult(1, 0)).WillReturnError(err)
		}
		output := s.Set(value.input)
		if output != value.expectedOutput {
			t.Errorf("failed  output %v test case failed %v", output, i+1)
		}

	}

}

func TestGet(t *testing.T) {

	testcases := []struct {
		desc           string
		id             int
		expectedOutput Car
	}{

		{"valid value", 4, Car{4, "XUV", "MGN2", "Petrol"}},
		{"valid value", 1, Car{1, "Safari", "MGN1", "Petrol"}},
		{"valid value", 2, Car{2, "Wagner", "FHG", "Petrol"}},
	}
	var s Store
	var err error
	db, mock, err := sqlmock.New()
	s.db = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	for i, value := range testcases {

		rows := sqlmock.NewRows([]string{"Id", "Name", "Model", "EngineType"}).
			AddRow(value.expectedOutput.Id, value.expectedOutput.Name, value.expectedOutput.Model, value.expectedOutput.EngineType)
		mock.ExpectQuery("select (.+) from Car where id=?").
			WithArgs(value.id).WillReturnRows(rows).WillReturnError(err)
		output := s.Get(value.id)
		if output != value.expectedOutput {
			t.Errorf("failed  output %v test case failed %v", output, i+1)
		}

	}

}
func TestDelete(t *testing.T) {

	testcases := []struct {
		desc           string
		id             int
		expectedOutput bool
	}{
		{"valid values", 2, true},
		{"NOT PRESENT IN DATABASE", 6, true},
	}
	var s Store
	var err error
	db, mock, err := sqlmock.New()
	s.db = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// now we execute our method
	for i, value := range testcases {

		mock.ExpectExec("DELETE FROM Car WHERE id").
			WithArgs(value.id).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(err)
		output := s.Delete(value.id)
		if output != value.expectedOutput {
			t.Errorf("failed  output %v test case failed %v", output, i+1)
		}

	}

}
