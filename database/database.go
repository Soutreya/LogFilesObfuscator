package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Order struct {
	FileID       int    `json:"file_id"`
	OrderID      string `json:"order_id"`
	Date         string `json:"date"`
	CustomerID   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	PizzaID      string `json:"pizza_id"`
	PizzaName    string `json:"pizza_name"`
	CardNum      string `json:"card_num"`
	Cost         int    `json:"cost"`
}

func Save(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("File opened successfully")
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("File read successfully")
	var ord []Order
	json.Unmarshal(byteValue, &ord)
	db, err := sql.Open("mysql", "root:Lead82_raz@tcp(127.0.0.1:3306)/log")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database opened successfully")
	defer db.Close()
	for i := 0; i < len(ord); i++ {
		query := "INSERT INTO orders VALUES(" + strconv.Itoa(ord[i].FileID) + ", '" + ord[i].OrderID + "', '" + ord[i].Date + "', '" + ord[i].CustomerID + "', '" + ord[i].CustomerName + "', '" + ord[i].PizzaID + "', '" + ord[i].PizzaName + "', '" + ord[i].CardNum + "', " + strconv.Itoa(ord[i].Cost) + ")"
		_, err := db.Query(query)
		if err != nil {
			panic(err)
		}
	}
}
