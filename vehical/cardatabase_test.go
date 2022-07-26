package vehical

import (
	"databasesample/driver"
	"fmt"
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
		{"valid values", Car{4, "XUV", "MGN2", "Petrol"}, true},
		{"duplicate not allowed", Car{4, "XUV2", "MGN3", "Petrol"}, false},
	}
	var s Store
	var err error
	myconfig := driver.MysqlConfig{"root", "Ra1190@cm", "3306", "localhost", "test"}
	s.db, err = driver.ConnectMysql(myconfig)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	for i, value := range testcases {
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
		{"negative id`s not allowed", -4, Car{0, "", "", ""}},
		{"valid value", 1, Car{1, "Safari", "MGN1", "Petrol"}},
		{"valid value", 2, Car{2, "Wagner", "FHG", "Petrol"}},
	}
	var s Store
	var err error
	myconfig := driver.MysqlConfig{"root", "Ra1190@cm", "3306", "localhost", "test"}
	s.db, err = driver.ConnectMysql(myconfig)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	for i, value := range testcases {
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
		{"NOT PRESENT IN DATABASE", 6, false},
	}
	var s Store
	var err error
	myConfig := driver.MysqlConfig{User: "root", Pass: "Ra1190@cm", Port: "3306", Host: "localhost", Dbname: "test"}
	s.db, err = driver.ConnectMysql(myConfig)
	if err == nil {
	}
	for i, value := range testcases {
		output := s.Delete(value.id)
		if output != value.expectedOutput {
			t.Errorf("failed  output %v test case failed %v", output, i+1)
		}
	}
}
