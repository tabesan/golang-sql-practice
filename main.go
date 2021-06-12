package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	queryInsertItem = "INSERT INTO items(item, price) VALUES (?, ?);"
)

type goods struct {
	name  string
	price int
}

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "tabe:natou@/practice")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO items (item, price) VALUES (?, ?)`)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("ball", 200)
	if err != nil {
		fmt.Println(err)
	}

	items, err := db.Query("SELECT * FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer items.Close()

	for items.Next() {
		var (
			item  string
			price int
		)

		if err := items.Scan(&item, &price); err != nil {
			log.Fatal(err)
		}
		fmt.Println("item:", item, "price:", price)
	}

	_, err = db.Query("DELETE FROM items WHERE item=?", "ball")
	if err != nil {
		log.Fatal(err)
	}
}
