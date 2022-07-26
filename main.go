package main

import (
	"databasesample/driver"
	"databasesample/vehical"
	_ "databasesample/vehical"
	"fmt"
	_ "fmt"
)

func main() {

	var s vehical.Store
	var err error
	myconfig := driver.MysqlConfig{"root", "Ra1190@cm", "3306", "localhost", "test"}
	s.Db, err = driver.ConnectMysql(myconfig)

	var c vehical.Car = vehical.Car{234, "Safari", "MGN1", "Petrol"}
	if err != nil {

		fmt.Errorf("%v", err)
	}
	// calling set function
	output := s.Set(c)
	fmt.Println(output)
	var car vehical.Car
	//calling get function
	car = s.Get(234)
	fmt.Printf("Id:%v\nName:%v\nModel:%v\nEngineType:%v\n", car.Id, car.Name, car.Model, car.EngineType)

	//calling delete function

	output1 := s.Delete(234)

	fmt.Println(output1)

}
