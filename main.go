package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	jmDictPath  = "./yomichanJsonInput/jmdict_english/"
	hanvietDict = "./SQLiteDB/finalDBSHanVietSound_Production.sqlite"
	kanjiQuery  = `SELECT field2, field3 FROM hanviet where field1=?`
	// query2      = `SELECT field2 FROM meaning where field1=?`
)

type EntryType int

const (
	String EntryType = iota
	Number
	StringSlice
)

// ["卵泡立て器","たまごあわだてき","n","",2,["egg beater"],1549320,""]
type Entry interface {
	Type() EntryType
}

func main() {
	db, err := sql.Open("sqlite3", hanvietDict)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	files, err := ioutil.ReadDir(jmDictPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !strings.Contains(file.Name(), "json") ||
			file.Name() == "index.json" ||
			file.Name() == "tag_bank_1.json" {
			continue
		}
		file, err := os.Open(jmDictPath + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		decoder := json.NewDecoder(file)
		if _, err := decoder.Token(); err != nil {
			log.Fatal(err)
		}
		for decoder.More() {

		}
	}

	rows, err := db.Query(kanjiQuery, "雨子")
	if err != nil {
		log.Fatal(err)
	}

	readRow(rows)
}

func readRow(rows *sql.Rows) {

	defer rows.Close()
	counter := 0
	for rows.Next() {
		var field2, field3 string
		if err := rows.Scan(&field2, &field3); err != nil {
			log.Fatal(err)
		}
		fmt.Println(field2, field3)
		counter++
	}

	log.Printf("done, got %d results", counter)
}
