package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int    `json:"id,string,omitempty"`
	Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage...")
	fmt.Fprintf(w, "Welcome to the HomePage.\n")
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllUser...")

	// Connecting to mySQL Database...
	db := dbConn()
	defer db.Close()

	query := `
		select * from user;
	`
	rows, err := db.Query(query)

	user := User{}
	res := []User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		//	fmt.Println("ID: ", user.Id, " Name: ", user.Name)
		res = append(res, user)
	}

	// Closing DB connnection...

	fmt.Fprintf(w, "Welcome to the getAllUser.\n")
	json.NewEncoder(w).Encode(res)
}

func getSingleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getSingleUser...")
	vars := mux.Vars(r)
	key := vars["id"]

	// Connecting to mySQL Database...
	db := dbConn()

	query := "select * from user where id=" + key + ";"
	rows, err := db.Query(query)

	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			panic(err.Error())
		}
	}

	// Closing DB connnection...
	defer db.Close()

	fmt.Fprintf(w, "Welcome to the getSingleUser.\n")
	json.NewEncoder(w).Encode(user)
}

func addNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: addNewUser...")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	fmt.Println(user)
	dbInsert(user)
	fmt.Fprintf(w, "Welcome to the addNewUser.\n")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateUser...")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	dbUpdate(user)
	fmt.Fprintf(w, "Welcome to the updateUser.\n")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteUser...")
	//reqBody, _ := ioutil.ReadAll(r.Body)
	id := mux.Vars(r)["id"]

	i, err := strconv.Atoi(id)
	errorCheck(err)
	n_rows_affected := dbDelete(int(i))

	fmt.Fprintf(w, "Welcome to the deleteUser.\n")
	fmt.Fprintf(w, fmt.Sprint("Number of rows affected: ", n_rows_affected))
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// HomePage
	myRouter.HandleFunc("/", homePage)

	// Show all users
	myRouter.HandleFunc("/show", getAllUser)

	// Show users with given Id
	myRouter.HandleFunc("/show/{id}", getSingleUser)

	// Add user
	myRouter.HandleFunc("/add", addNewUser).Methods("POST")

	// Update
	myRouter.HandleFunc("/update", updateUser).Methods("PUT")

	// Delete
	myRouter.HandleFunc("/delete/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("CRUD using mySQL and Golang...")
	handleRequest()
}
