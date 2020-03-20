package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"waitformore"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5234, "wait4more", "zCpzQ7VF7%GFmTM", "wait4more")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	repository := waitformore.PostgresqlQueueRepository{DatabaseConnectionPool: *db}
	definition, repositoryError := repository.CreateNewQueue("asdf")
	if repositoryError != nil {
		panic(repositoryError)
	}
	fmt.Printf("%v\n", definition)

}
