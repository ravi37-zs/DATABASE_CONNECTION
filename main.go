package main

import (
	"databasesample/driver"
	"fmt"
)

func main() {

	myconfig := driver.MysqlConfig{"root", "Ra1190@cm", "3306", "localhost", "test"}

	db, err := driver.ConnectMysql(myconfig)
	fmt.Println(db, err)
<<<<<<< HEAD
	fmt.Println("this is ravi1")
=======

}
