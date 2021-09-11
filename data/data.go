package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS cards (
		"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Golingo table created")
}

func InsertCard(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO cards(word, definition, category) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted word successfully")
}

func DisplayAllCards() {
	row, err := db.Query("SELECT * FROM cards ORDER BY word")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string
		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "] ", word, "—", definition)
	}
}
