package main

import (
	"database/sql"
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
	limitData = 10
)
const (
	select_all_stmt  = "SELECT * FROM Txn"
	select_user_stmt = "SELECT * FROM Txn WHERE trim(userName)="
	select_bank_stmt = "SELECT * FROM Txn WHERE trim(bankName)="
	insert_stmt      = "INSERT INTO Txn(timeStamp, txnId, upiId, userName, bankName, txnAmount) VALUES(?, ?, ?, ?, ?, ?)"
	delete_stmt      = "DELETE FROM Txn WHERE trim(userName)=?"
	update_stmt      = "UPDATE Txn SET name=? WHERE id=?"
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

func queryExecutor(query string) []DashboardResult {
	db := dbConn()
	defer db.Close()
	rows, _ := db.Query(query)

	txn := DashboardResult{}
	res := []DashboardResult{}
	for rows.Next() {
		err := rows.Scan(&txn.Name, &txn.Value)
		errorCheck(err)
		res = append(res, txn)
	}
	return res
}

func dashboardData() ([][]DashboardResult, Resp) {

	query := "SELECT userName, SUM(txnAmount) FROM Txn GROUP BY userName ORDER BY SUM(txnAmount) DESC limit 20;"
	res1 := queryExecutor(query)
	query = "SELECT bankName, SUM(txnAmount) FROM Txn GROUP BY bankName ORDER BY SUM(txnAmount) DESC limit 20;"
	res2 := queryExecutor(query)
	query = "SELECT userName, SUM(txnAmount) FROM Txn GROUP BY userName ORDER BY SUM(txnAmount) ASC limit 5;"
	res3 := queryExecutor(query)
	query = "SELECT bankName, SUM(txnAmount) FROM Txn GROUP BY bankName ORDER BY SUM(txnAmount) ASC limit 5;"
	res4 := queryExecutor(query)

	query = "SELECT SUM(txnAmount) FROM Txn;"
	db := dbConn()
	defer db.Close()
	rows, _ := db.Query(query)

	txn := Resp{}
	for rows.Next() {
		err := rows.Scan(&txn.NumRowsAffected)
		errorCheck(err)
	}

	return [][]DashboardResult{res1, res2, res1[:8], res2[:8], res3, res4}, txn
}

func getAllData() []Transaction {
	db := dbConn()
	defer db.Close()

	rows, _ := db.Query(select_all_stmt)

	txn := Transaction{}
	res := []Transaction{}
	i := 0
	for rows.Next() {
		err := rows.Scan(&txn.TimeStamp, &txn.TxnId, &txn.UpiId, &txn.UserName, &txn.BankName, &txn.TxnAmount)
		errorCheck(err)
		res = append(res, txn)
		if i == limitData {
			break
		}
		i++
	}
	return res
}

func getUserData(userName string) []Transaction {
	db := dbConn()
	defer db.Close()

	query := select_user_stmt + "'" + userName + "';"
	rows, err := db.Query(query)
	errorCheck(err)

	txn := Transaction{}
	res := []Transaction{}
	for rows.Next() {
		err := rows.Scan(&txn.TimeStamp, &txn.TxnId, &txn.UpiId, &txn.UserName, &txn.BankName, &txn.TxnAmount)
		errorCheck(err)
		res = append(res, txn)
	}
	return res
}

func getBankData(bankName string) []Transaction {
	db := dbConn()
	defer db.Close()

	query := select_bank_stmt + "'" + bankName + "';"
	rows, err := db.Query(query)
	errorCheck(err)

	txn := Transaction{}
	res := []Transaction{}
	for rows.Next() {
		err := rows.Scan(&txn.TimeStamp, &txn.TxnId, &txn.UpiId, &txn.UserName, &txn.BankName, &txn.TxnAmount)
		errorCheck(err)
		res = append(res, txn)
	}
	return res
}

func addTransaction(txn Transaction) int {
	db := dbConn()
	defer db.Close()
	stmt, err := db.Prepare(insert_stmt)
	errorCheck(err)
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	// fmt.Println("Data added:", txn.TimeStamp, txn.TxnId, txn.UpiId, txn.UserName, txn.BankName, txn.TxnAmount)
	res, err := stmt.Exec(txn.TimeStamp, txn.TxnId, txn.UpiId, txn.UserName, txn.BankName, txn.TxnAmount)
	errorCheck(err)

	n_rows_affected, err := res.RowsAffected()
	errorCheck(err)

	return int(n_rows_affected)
}

func deleteUserData(userName string) int {
	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare(delete_stmt)
	errorCheck(err)

	res, err := stmt.Exec(userName)
	errorCheck(err)

	n_rows_affected, err := res.RowsAffected()
	errorCheck(err)

	defer db.Close()
	defer stmt.Close()
	return int(n_rows_affected)
}

// func main() {
// 	// res := getAllData()
// 	// fmt.Println(res)
// 	// fmt.Println("####################################")

// 	fmt.Println(getUserData("user15"))
// 	fmt.Println("####################################")
// 	// fmt.Println(getBankData("RBL Bank"))
// 	txn := Transaction{"2021-05-28 04:20:12", "3IOUGWHsUYklx9MDgldM", "user15@YES", "user15", "YES Bank", 5272}
// 	fmt.Println(addTransaction(txn))
// 	// fmt.Println("####################################")
// 	// fmt.Println(deleteUserData("user15"))

// 	fmt.Println("####################################")
// 	fmt.Println(getUserData("user15"))
// }
