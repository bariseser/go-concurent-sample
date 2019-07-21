package main

import (
	"database/sql"
	"fmt"
	"go-concurrent-sample/db"
	"go-concurrent-sample/models"
	"log"
	"math"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := db.InitDB()

	if err != nil {
		log.Panic(err)
	}

	env := &Env{db: db}

	var totalCount int
	queryError := db.QueryRow("SELECT count(*) as total from users").Scan(&totalCount)

	if queryError != nil {
		panic(queryError.Error())
	}

	fmt.Println("Total Users:", totalCount)

	limit := 10
	fmt.Println("Limit:", limit)
	pages := round(limit, totalCount)
	fmt.Println("Total Concurrent:", pages, "\n")

	for i := 0; i < pages; i++ {
		fmt.Println("Concurrent:\t", i)
		go query(env, i*limit, limit)
	}

	var userInput float64
	fmt.Scanln(&userInput)
}

func round(firstNumber int, secondNumber int) int {
	return int(math.Ceil(float64(secondNumber) / float64(firstNumber)))
}

func query(env *Env, start int, end int) {

	sql := fmt.Sprintf("SELECT id, fullname FROM users limit %d,%d", start, end)

	queryResult, err := models.GetUser(env.db, sql)

	if err != nil {
		panic(err)
	}

	for _, row := range queryResult {
		fmt.Printf("%s, %s\n", row.ID, row.Fullname)
	}
}
