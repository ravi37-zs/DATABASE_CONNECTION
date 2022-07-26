package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	User   string
	Pass   string
	Port   string
	Host   string
	Dbname string
}

func ConnectMysql(myconfig MysqlConfig) (*sql.DB, error) {
	constr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", myconfig.User, myconfig.Pass, myconfig.Host, myconfig.Port, myconfig.Dbname)
	fmt.Println(constr)
	db, err := sql.Open("mysql", constr)
	if err != nil {
		fmt.Println("err is: ", err)
		return nil, err
	}

	db.Exec("create table if not exists Car(id int PRIMARY KEY,name varchar(255),Model varchar(255),EngineType varchar(255));")

	return db, err
}
