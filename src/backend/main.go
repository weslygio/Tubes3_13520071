package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/weslygio/Tubes3_13520071/src/backend/pkg/dna"
	"github.com/weslygio/Tubes3_13520071/src/backend/pkg/timeconv"
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
	Tgl          string  `json:"tanggal"`
	NamaPasien   string  `json:"namaPasien"`
	NamaPenyakit string  `json:"namaPenyakit"`
	Kemiripan    float64 `json:"kemiripan"`
	Hasil        bool    `json:"hasil"`
}

type Penyakit struct {
	NamaPenyakit string `json:"namaPenyakit"`
	DNASequence  string `json:"dnaSequence"`
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
	}))

	router.GET("/logs/", getLogs)
	router.GET("/logs/:query", getLogs)
	router.POST("/logs", postLogs)
	router.POST("/diseases", postDisease)

	router.Run("localhost:8080")
}

func postDisease(c *gin.Context) {
	var penyakit Penyakit

	// Bind request body
	err := c.BindJSON(&penyakit)
	if err != nil {
		panic(err.Error())
	}

	// Sanitize input
	if strings.Trim(penyakit.NamaPenyakit, " ") == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	isDNAValid := dna.IsDNAValid(penyakit.DNASequence)
	if !isDNAValid {
		c.Status(http.StatusNotAcceptable)
		return
	}

	// Insert into database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO penyakit(namaPenyakit,sequence) VALUES(? , ?)", penyakit.NamaPenyakit, penyakit.DNASequence)
	if err != nil {
		c.Status(http.StatusConflict)
		return
	}
	defer insert.Close()

	c.Status(http.StatusOK)
}

func postLogs(c *gin.Context) {
	var pasien Pasien
	var logPasien LogPasien
	var disease Penyakit

	// Bind request body
	err := c.BindJSON(&pasien)
	if err != nil {
		panic(err.Error())
	}

	// Sanitize input
	isValid := dna.IsDNAValid(pasien.DNASequence)
	if !isValid {
		c.Status(http.StatusNotAcceptable)
		return
	}

	// Find disease in database
	disease = getDiseasebyName(pasien.NamaPenyakit)
	if disease.DNASequence == "" {
		c.Status(http.StatusFailedDependency)
		return
	}

	// Compare DNA and add to log
	match, similarity := dna.IsDNAMatched(pasien.DNASequence, disease.DNASequence)
	logPasien.Tgl = timeconv.DateToString(time.Now().Date())
	logPasien.NamaPasien = pasien.NamaPasien
	logPasien.NamaPenyakit = pasien.NamaPenyakit
	logPasien.Kemiripan = float64(similarity)
	logPasien.Hasil = match

	// Insert into database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO logPasien(tglCheckUp, namaPasien, namaPenyakit, kemiripan, hasil) VALUES(?,?,?,?,?)",
		timeconv.DateToYYYYMMDD(time.Now().Date()), logPasien.NamaPasien, logPasien.NamaPenyakit, logPasien.Kemiripan, logPasien.Hasil)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	c.JSON(http.StatusOK, logPasien)
}

func getLogs(c *gin.Context) {
	// Parse query
	valid, time, disease := dna.ParsePrediction(c.Param("query"))
	if !valid {
		c.Status(http.StatusNotAcceptable)
		return
	}

	arr := []LogPasien{}
	var res *sql.Rows
	var temp string

	// Select matching data from database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if !time.IsZero() && disease == "" {
		// search by given date
		res, err = db.Query("SELECT * FROM logPasien WHERE tglCheckUp = ?", timeconv.DateToYYYYMMDD(time.Date()))
	} else if time.IsZero() && disease != "" {
		// search by given disease
		res, err = db.Query("SELECT * FROM logPasien WHERE LOWER(namaPenyakit) LIKE LOWER(?)", "%"+disease+"%")
	} else {
		// search by given date and disease
		res, err = db.Query("SELECT * FROM logPasien WHERE tglCheckUp = ? AND LOWER(namaPenyakit) LIKE LOWER(?)", timeconv.DateToYYYYMMDD(time.Date()), "%"+disease+"%")
	}

	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	// Append all matching logs to response body
	for res.Next() {
		var logPasien LogPasien
		err := res.Scan(&temp, &logPasien.Tgl, &logPasien.NamaPasien, &logPasien.NamaPenyakit, &logPasien.Kemiripan, &logPasien.Hasil)
		if err != nil {
			log.Fatal(err)
		} else {
			arr = append(arr, logPasien)
		}
	}

	c.JSON(http.StatusOK, arr)
}

func getDiseasebyName(name string) Penyakit {
	var disease Penyakit
	var temp string

	// Select disease on name
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
