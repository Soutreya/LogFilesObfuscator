package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
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
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/log")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database opened successfully")
	defer db.Close()
	var wg sync.WaitGroup
	for i := 0; i < len(ord); i++ {
		wg.Add(1)
		go func(data Order) {
			query := "INSERT INTO orders VALUES(" + strconv.Itoa(data.FileID) + ", '" + data.OrderID + "', '" + data.Date + "', '" + data.CustomerID + "', '" + data.CustomerName + "', '" + data.PizzaID + "', '" + data.PizzaName + "', '" + data.CardNum + "', " + strconv.Itoa(data.Cost) + ")"
			_, err := db.Query(query)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(ord[i])
	}
	wg.Wait()
}
