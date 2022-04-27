package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user = "YJk8zu10qJ"
	db_pass = "b3b50r3cmf"
	db_addr = "remotemysql.com"
	db_name = "YJk8zu10qJ"
)

type LogPasien struct {
	IdPasien     string
	Tgl          string
	NamaPasien   string
	NamaPenyakit string
	Kemiripan    int
	Hasil        bool
}

type Penyakit struct {
	IdPenyakit   string
	NamaPenyakit string
	DnaSequence  string
}

func insertPenyakit(namaPenyakit string, sequence string) {
	// open database
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pass, db_addr, db_name)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO penyakit(namaPenyakit,sequence) VALUES(? , ?)", namaPenyakit, sequence)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func insertLogPasien(tgl string, namaPasien string, namaPenyakit string, kemiripan int, hasil bool) {
	// open database
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pass, db_addr, db_name)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO logPasien(tglCheckUp, namaPasien, namaPenyakit, kemiripan, hasil) VALUES(?,?,?,?)",
		tgl, namaPasien, namaPenyakit, kemiripan, hasil)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func getLogbyDate(date string) []string {
	// open database
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pass, db_addr, db_name)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM logPasien WHERE tglCheckUp = ?", date)
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	// temporary variable
	arr := []string{}

	if res.Next() {
		var logPasien LogPasien
		err := res.Scan(&logPasien.IdPasien, &logPasien.Tgl, &logPasien.NamaPasien, &logPasien.NamaPenyakit, &logPasien.Kemiripan, &logPasien.Hasil)
		if err != nil {
			log.Fatal(err)
		} else {
			kemiripan := fmt.Sprintf("%d", logPasien.Kemiripan)
			hasil := fmt.Sprintf("%v", logPasien.Hasil)
			var temp = logPasien.Tgl + "-" + logPasien.NamaPasien + "-" + logPasien.NamaPenyakit + "-" + kemiripan + "-" + hasil
			arr = append(arr, temp)
		}

	}
	return arr
}

func getLogbyDisease(disease string) []string {
	// open database
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pass, db_addr, db_name)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM logPasien WHERE namaPenyakit = ?", disease)
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	// temporary variable
	arr := []string{}

	if res.Next() {
		var logPasien LogPasien
		err := res.Scan(&logPasien.IdPasien, &logPasien.Tgl, &logPasien.NamaPasien, &logPasien.NamaPenyakit, &logPasien.Kemiripan, &logPasien.Hasil)
		if err != nil {
			log.Fatal(err)
		} else {
			kemiripan := fmt.Sprintf("%d", logPasien.Kemiripan)
			hasil := fmt.Sprintf("%v", logPasien.Hasil)
			var temp = logPasien.Tgl + "-" + logPasien.NamaPasien + "-" + logPasien.NamaPenyakit + "-" + kemiripan + "-" + hasil
			arr = append(arr, temp)
		}

	}
	return arr
}

func getLogbyDateAndDisease(date string, disease string) []string {
	// open database
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pass, db_addr, db_name)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM logPasien WHERE tglCheckUp = ? AND namaPenyakit = ?", date, disease)
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	// temporary variable
	arr := []string{}

	if res.Next() {
		var logPasien LogPasien
		err := res.Scan(&logPasien.IdPasien, &logPasien.Tgl, &logPasien.NamaPasien, &logPasien.NamaPenyakit, &logPasien.Kemiripan, &logPasien.Hasil)
		if err != nil {
			log.Fatal(err)
		} else {
			kemiripan := fmt.Sprintf("%d", logPasien.Kemiripan)
			hasil := fmt.Sprintf("%v", logPasien.Hasil)
			var temp = logPasien.Tgl + "-" + logPasien.NamaPasien + "-" + logPasien.NamaPenyakit + "-" + kemiripan + "-" + hasil
			arr = append(arr, temp)
		}

	}
	return arr
}
