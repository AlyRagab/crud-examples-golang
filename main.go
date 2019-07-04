package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

var err error
var db *sql.DB

func main() {
	// Database Connection
	db, err = sql.Open("mysql", "root:pass@tcp(localhost:3306)/go")
	check(err)
	defer db.Close()

	// Handle CRUD Functions
	fmt.Println("Connected Successfully !!")
	http.HandleFunc("/", index)
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/get", get)
	http.HandleFunc("/create", create)
	http.HandleFunc("/createdb", createdb)
	http.HandleFunc("/insert", insert)
	check(err)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "Connected Successfully !!")
	check(err)
}

// Drop Database Function
func drop(w http.ResponseWriter, req *http.Request) {
	dr, err := db.Prepare(`DROP DATABASE go;`)
	check(err)
	_, err = dr.Exec()
	check(err)

	fmt.Fprint(w, "Done!!")
}

func get(w http.ResponseWriter, r *http.Request) {
	sel, err := db.Query(`SELECT * FROM go.users;`)
	check(err)

	var s, name string
	s = "Data is : \n"

	for sel.Next() {
		err = sel.Scan(&name)
		check(err)
		s += name + "\n"
	}

	fmt.Fprint(w, s)
}
func create(w http.ResponseWriter, r *http.Request) {
	crt, err := db.Prepare(`CREATE TABLE users (name VARCHAR(20))`)
	check(err)

	c, err := crt.Exec()
	check(err)

	n, err := c.RowsAffected()
	check(err)

	fmt.Fprint(w, n)
}
func createdb(w http.ResponseWriter, r *http.Request) {
	crtdb, err := db.Prepare(`CREATE DATABASE go;`)
	check(err)

	_, err = crtdb.Exec()
	check(err)

	fmt.Fprint(w, "Created !!")
}
func insert(w http.ResponseWriter, r *http.Request) {
	ins, err := db.Prepare(`INSERT INTO users VALUES ('Aly');`)
	check(err)

	i, err := ins.Exec()
	check(err)

	n, err := i.RowsAffected()
	check(err)

	fmt.Fprint(w, "Inserted !", n)
}
