package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/xshyamx/h2go"
)

func must(ctx string, err error) {
	if err != nil {
		log.Fatalf("%s - error : %s", ctx, err)
	}
}

type record struct {
	id   int
	name string
	age  int
}

func main() {
	log.Printf("H2GO Example")
	// ?mem=true&logging=info
	conn, err := sql.Open("h2", "h2://sa@localhost/test")
	must("Open connection", err)
	defer conn.Close()
	// Drop table if exists
	log.Printf("DROP TABLE")
	_, err = conn.Exec("DROP TABLE test IF EXISTS")
	must("Drop table", err)
	// Create table
	log.Printf("CREATE TABLE")
	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS test (id int auto_increment, name varchar(50), age int)")
	must("Create table", err)
	// Insert
	log.Printf("Insert")
	stmt, err := conn.Prepare("INSERT INTO test(name, age) VALUES(?, ?)")
	must("Prepare statement", err)
	for i, name := range []string{"John", "Jane", "Jack", "Jill"} {
		_, err := stmt.Exec(name, i+30)
		must(fmt.Sprintf("Inserting %s", name), err)
	}
	{
		var r record
		// Statment Query
		log.Printf("Statement Query")
		stmt, err := conn.Prepare("SELECT * FROM test where name=?")
		must("Prepare stmt", err)
		stmt.QueryRow("Jill").Scan(&r.id, &r.name, &r.age)
		log.Printf("Record: %+v", r)
	}
	// Query
	log.Printf("Query")
	var r record
	err = conn.QueryRow("SELECT * FROM test where name=?", "Jill").Scan(&r.id, &r.name, &r.age)
	must("Finding Jill", err)
	log.Printf("Record: %+v", r)
	// Update
	log.Printf("Update")
	_, err = conn.Exec("UPDATE test SET name = 'Juan' WHERE id = ?", 1)
	must("Update 1", err)
	// Delete
	log.Printf("Delete")
	_, err = conn.Exec("DELETE FROM test WHERE id = ?", 1)
	must("Delete 1", err)
	time.Sleep(5 * time.Second)
	log.Printf("Done")
}
