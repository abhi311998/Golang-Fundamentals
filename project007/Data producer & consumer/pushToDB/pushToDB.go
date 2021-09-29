package pushToDB

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	
	"github.com/abhi311998/txnDataGen"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "facebook"
	dbName   = "myTestDB"
)
const (
	insert_stmt = "INSERT INTO Txn(timeStamp, txnId, upiId, userName, bankName, txnAmount) VALUES(?, ?, ?, ?, ?, ?)"
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

func AddTransaction(txn txnDataGen.Transaction) (int, error) {
	db := dbConn()
	defer db.Close()
	stmt, err := db.Prepare(insert_stmt)
	errorCheck(err)
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	res, err := stmt.Exec(txn.TimeStamp, txn.TxnId, txn.UpiId, txn.UserName, txn.BankName, txn.TxnAmount)
	errorCheck(err)

	n_rows_affected, err := res.RowsAffected()
	errorCheck(err)

	return int(n_rows_affected), err
}