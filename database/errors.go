package database

import (
	"database/sql"
	"fmt"
	"log"
)

func HandleError(message string, err error) {
	switch err {
	case sql.ErrNoRows:
		fmt.Println("There is no data")
	case nil:
		fmt.Println("No Error")
	default:
		log.Fatalf(message+": %v", err)
	}
}
