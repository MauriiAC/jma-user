package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func DbConnect() error {
	databaseURL := "root:H2Onteyo@tcp(jma-ddbb.cawwunnevkkt.us-east-1.rds.amazonaws.com:3306)/jma-ddbb"
	Db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer Db.Close()

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
