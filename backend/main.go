package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/weslygio/Tubes3_13520071/src/pkg/dna"
	"github.com/weslygio/Tubes3_13520071/src/pkg/timeconv"
)

const (
	db_user = "YJk8zu10qJ"
	db_pass = "b3b50r3cmf"
	db_addr = "remotemysql.com"
	db_name = "YJk8zu10qJ"
)

var dataSourceName string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pass, db_addr, db_name)

type Pasien struct {
	NamaPasien   string `json:"namaPasien"`
	DNASequence  string `json:"dnaSequence"`
	NamaPenyakit string `json:"namaPenyakit"`
}

type LogPasien struct {
	Tgl          string `json:"title"`
	NamaPasien   string `json:"namaPasien"`
	NamaPenyakit string `json:"namaPenyakit"`
	Kemiripan    int    `json:"kemiripan"`
	Hasil        bool   `json:"hasil"`
}

type Penyakit struct {
	NamaPenyakit string `json:"namaPenyakit"`
	DNASequence  string `json:"dnaSequence"`
}

func main() {
	router := gin.Default()
	router.GET("/logs?query=:query", getLogs)
	router.POST("/logs", postLogs)
	router.POST("/diseases", postDisease)

	router.Run("localhost:8080")
}

func postDisease(c *gin.Context) {
	var requestDisease Penyakit
	err := c.BindJSON(&requestDisease)
	if err != nil {
		panic(err.Error())
	}

	// open database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert new disease
	insert, err := db.Query("INSERT INTO penyakit(namaPenyakit,sequence) VALUES(? , ?)", requestDisease.NamaPenyakit, requestDisease.DNASequence)
	if err != nil {
		// handle double disease
		return
	}
	defer insert.Close()
}

func postLogs(c *gin.Context) {
	var pasien Pasien
	var logPasien LogPasien
	var disease Penyakit

	err := c.BindJSON(&pasien)
	if err != nil {
		panic(err.Error())
	}

	// open database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//proses
	isValid := dna.IsDNAValid(pasien.DNASequence)
	if !isValid {
		// handle invalid DNA
		return
	}

	disease = getDiseasebyName(pasien.NamaPenyakit)
	if disease.DNASequence == "" {
		// handle when there's no such disease
		return
	}
	match, similarity := dna.IsDNAMatched(pasien.DNASequence, disease.DNASequence)

	// set up logPasien
	logPasien.Tgl = timeconv.DateToString(time.Now().Date())
	logPasien.NamaPasien = pasien.NamaPasien
	logPasien.NamaPenyakit = pasien.NamaPenyakit
	logPasien.Kemiripan = int(similarity)
	logPasien.Hasil = match

	// insert new log
	insert, err := db.Query("INSERT INTO logPasien(tglCheckUp, namaPasien, namaPenyakit, kemiripan, hasil) VALUES(?,?,?,?)",
		logPasien.Tgl, logPasien.NamaPasien, logPasien.NamaPenyakit, logPasien.Kemiripan, logPasien.Hasil)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func getLogs(c *gin.Context) {
	valid, date, disease := dna.ParsePrediction(c.Param("query"))
	if !valid {
		return
	}

	// open database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// temporary variable
	arr := []LogPasien{}
	var res *sql.Rows
	var temp string

	if !date.IsZero() && disease == "" {
		// search by given date
		res, err = db.Query("SELECT * FROM logPasien WHERE tglCheckUp = ?", date)
	} else if date.IsZero() && disease != "" {
		// search by given disease
		res, err = db.Query("SELECT * FROM logPasien WHERE namaPenyakit = ?", disease)
	} else {
		// search by given date and disease
		res, err = db.Query("SELECT * FROM logPasien WHERE tglCheckUp = ? AND namaPenyakit = ?", date, disease)
	}
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	for res.Next() {
		var logPasien LogPasien
		err := res.Scan(&temp, &logPasien.Tgl, &logPasien.NamaPasien, &logPasien.NamaPenyakit, &logPasien.Kemiripan, &logPasien.Hasil)
		if err != nil {
			log.Fatal(err)
		} else {
			arr = append(arr, logPasien)
		}
	}

	// get request
	c.IndentedJSON(http.StatusOK, arr)
}

func getDiseasebyName(name string) Penyakit {
	// temporary variable
	var disease Penyakit
	var temp string

	// open database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res := db.QueryRow("SELECT * FROM penyakit WHERE namaPenyakit = ?", name)
	err = res.Scan(&temp, &disease.NamaPenyakit, &disease.DNASequence)
	if err == sql.ErrNoRows {
		disease.NamaPenyakit = ""
		disease.DNASequence = ""
	}

	return disease
}
