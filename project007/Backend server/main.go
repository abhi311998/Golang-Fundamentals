package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Transaction struct {
	TimeStamp string `json:"timeStamp"`
	TxnId     string `json:"txnId"`
	UpiId     string `json:"upiId"`
	UserName  string `json:"userName"`
	BankName  string `json:"bankName"`
	TxnAmount int64  `json:"txnAmount"`
}

type DashboardResult struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type Resp struct {
	NumRowsAffected int `json:"numRowsAffected"`
}

type DashboardResult2 struct {
	First  [][]DashboardResult `json:"first"`
	Second Resp                `json:"second"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage...")
	fmt.Fprintf(w, "Welcome to the HomePage.\n")
}

func getDashboardData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getDashboardData...", time.Now().Minute(), time.Now().Second())

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	result := DashboardResult2{}
	res, resp := dashboardData()
	result.First = res
	result.Second = resp

	json.NewEncoder(w).Encode(result)
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllUser...", time.Now().Minute(), time.Now().Second())

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	res := getAllData()

	json.NewEncoder(w).Encode(res)
}

func getSingleUserData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getSingleUserData...", time.Now().Minute(), time.Now().Second())

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	// fmt.Println(r.Header)
	// fmt.Println(r.Method)

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)
	key := vars["userName"]
	res := getUserData(key)

	json.NewEncoder(w).Encode(res)
}

func getSingleBankData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getSingleBankData...", time.Now().Minute(), time.Now().Second())

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)
	key := vars["bankName"]
	res := getBankData(key)

	json.NewEncoder(w).Encode(res)
}

func addNewTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: addNewTransaction...", time.Now().Minute(), time.Now().Second())

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println(r.Method)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var txn Transaction
	json.Unmarshal(reqBody, &txn)
	fmt.Printf("%T, %v\n", txn, txn)

	n_rows_affected := addTransaction(txn)

	fmt.Fprintf(w, fmt.Sprint("Number of rows affected: ", n_rows_affected))
}

func deleteUserTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteUserTransaction...", time.Now().Minute(), time.Now().Second())
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	// fmt.Println(r.Header)
	// fmt.Println(r.Method)

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)
	key := vars["userName"]

	resp := Resp{}
	resp.NumRowsAffected = deleteUserData(key)

	json.NewEncoder(w).Encode(resp)
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// HomePage
	myRouter.HandleFunc("/", homePage)

	// getDashboardData
	myRouter.HandleFunc("/dashboard", getDashboardData)

	// Get all txns getDashboardData
	myRouter.HandleFunc("/get", getAllUser)

	// Show txns of specific user
	myRouter.HandleFunc("/getUser/{userName}", getSingleUserData).Methods("GET")

	// Show txns of specific bank
	myRouter.HandleFunc("/getBank/{bankName}", getSingleBankData).Methods("GET")

	// Add a txn
	myRouter.HandleFunc("/addTxn", addNewTransaction)

	// // Delete
	myRouter.HandleFunc("/delete/{userName}", deleteUserTransaction)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("CRUD using mySQL and Golang...")
	handleRequest()
}
