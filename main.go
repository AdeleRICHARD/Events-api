package main

import (
	"database/sql"
	"events-api/testdb"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	const (
		host     = "172.17.0.1"
		port     = "5432"
		user     = "user1"
		password = "adele"
		dbname   = "db"
	)

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	println(psqlconn)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("err" + err.Error())
	}
	defer db.Close()

	testdb.TestInsert(db)
	/*
		args := Args{
			conn: psqlconn,
			port: "9000",
		} */

	/* if err := Run(args); err != nil {
		log.Println(err)
	} */
	fmt.Println("Successfully connected!")
}
