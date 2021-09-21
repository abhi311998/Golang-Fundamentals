package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "facebook"
	dbName   = "myTestDB"
)

const (
	select_all_stmt      = "SELECT * FROM user"
	select_specific_stmt = "SELECT * FROM user WHERE id=?"
	insert_stmt          = "INSERT INTO user(id, name) VALUES(?, ?)"
	delete_stmt          = "DELETE FROM user WHERE id=?"
	update_stmt          = "UPDATE user SET name=? WHERE id=?"
)

func errorCheck(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func dbConn() (db *sql.DB) {
	dataSourceName := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName
	db, err := sql.Open(dbDriver, dataSourceName)
	errorCheck(err)
	return db
}

func dbInsert(user User) {
	db := dbConn()
	defer db.Close()
	stmt, err := db.Prepare(insert_stmt)
	errorCheck(err)
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	res, err := stmt.Exec(user.Id, user.Name)
	errorCheck(err)

	id, err := res.LastInsertId()
	errorCheck(err)

	fmt.Println("Last insert id", id)

}

func dbDelete(id int) int64 {
	db := dbConn()
	stmt, err := db.Prepare(delete_stmt)
	errorCheck(err)

	res, err := stmt.Exec(id)
	errorCheck(err)

	n_rows_affected, err := res.RowsAffected()
	errorCheck(err)

	fmt.Println("Number of rows affected: ", n_rows_affected)

	defer db.Close()
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	return n_rows_affected
}

func dbUpdate(user User) {
	db := dbConn()
	stmt, err := db.Prepare(update_stmt)
	errorCheck(err)

	res, err := stmt.Exec(user.Name, user.Id)
	errorCheck(err)

	n_rows_affected, err := res.RowsAffected()
	errorCheck(err)

	fmt.Println("Number of rows affected: ", n_rows_affected)

	defer db.Close()
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
}
