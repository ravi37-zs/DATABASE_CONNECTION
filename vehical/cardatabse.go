package vehical

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}
type Car struct {
	Id         int
	Name       string
	Model      string
	EngineType string
}

func (s Store) Get(id int) (c Car) {

	if id < 0 {
		return Car{}
	}
	row, err := s.db.Query("select * from Car where id=?", id)
	if err != nil {
		fmt.Errorf("%v", err)
		fmt.Printf("%v this is error in line 24", err)
	}
	row.Next()
	row.Scan(&c.Id, &c.Name, &c.Model, &c.EngineType)
	row.Close()
	return
}

func (s Store) Set(c Car) bool {
	res, err := s.db.Exec("insert into Car values(?,?,?,?)", c.Id, c.Name, c.Model, c.EngineType)
	//fmt.Println("helooooooooooooooooooooo")

	rows, _ := res.RowsAffected()
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

	if rows == 0 {
		return false
	}
	return true
}

func (s Store) Delete(Id int) bool {
	res, err := s.db.Exec("DELETE FROM Car WHERE id=?;", Id)
	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Errorf("%v", err)

	}
	if rows == 0 {
		return false
	}
	return true
}
