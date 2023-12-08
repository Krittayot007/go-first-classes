package repo

import (
	"database/sql"
	"fmt"
	"log"
)

func Database() *sql.DB {
	connectionString := "postgres://username:password@your-database-url:5432/your-database-name"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL!")

	/* // Example query
	   rows, err := db.Query("SELECT * FROM your_table")
	   if err != nil {
	   	log.Fatal(err)
	   }
	   defer rows.Close()

	   // Process the query results
	   for rows.Next() {
	   	var column1, column2 string
	   	err := rows.Scan(&column1, &column2)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	fmt.Printf("Column1: %s, Column2: %s\n", column1, column2)
	   }

	   // Handle any errors from iterating over rows
	   if err := rows.Err(); err != nil {
	   	log.Fatal(err)
	   } */

	return db
}
