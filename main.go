package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=bookings user=mark.ledger password=")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer conn.Close()
	updateRows(conn)
	log.Println("Connected to database")
	err = conn.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to ping database: %v\n", err))
	}
	log.Println("Pinged database")

	err = getAllRows(conn)
	if err != nil {
		log.Println("Error getAllRows: %s", err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		log.Println("Error querying users: %v\n", err)
	}

	defer rows.Close()
	var id, firstName, lastName, email string
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &email)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(fmt.Sprintf("id: %s, firstName: %s, lastName: %s, email: %s", id, firstName, lastName, email))
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}
	return nil
}

func insertRows(conn *sql.DB) error {
	query := `INSERT INTO users(first_name, last_name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err := conn.Exec(query, "Kit", "Ledger", "kit@example.com", "bobo")
	if err != nil {
		log.Fatal("Error inserting rows: %v\n", err)
	}
	return nil
}

func updateRows(conn *sql.DB) error {
	query := `UPDATE users set first_name= $1 where id=$2`
	_, err := conn.Exec(query, "Kit", 6)
	if err != nil {
		log.Fatal("Error inserting rows: %v\n", err)
	}
	return nil
}
