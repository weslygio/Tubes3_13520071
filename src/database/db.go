package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user = "YJk8zu10qJ"
	db_pass = "b3b50r3cmf"
	db_addr = "remotemysql.com"
	db_name = "YJk8zu10qJ"
)

type logPasien struct {
	idPasien     int
	cuYear       time.Time
	cuMonth      time.Time
	cuDate       time.Time
	namaPasien   string
	namaPenyakit string
	kemiripan    int
	hasil        bool
}

type penyakit struct {
	idPenyakit   int
	namaPenyakit string
	dnaSequence  string
}

func main() {
	fmt.Println("Go Mysql")
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pass, db_addr, db_name)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO penyakit(namaPenyakit,sequence) VALUES('tes','AGAGAGATC')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successfully insert to penyakit")

}
