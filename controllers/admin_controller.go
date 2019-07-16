package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	. "projects/log_files_obfuscator/database"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Admin(c *gin.Context) {
	key := c.Param("key")
	value := c.Param("value")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/log")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database opened successfully")
	defer db.Close()
	data, err := db.Query("SELECT * FROM orders")
	if err != nil {
		panic(err)
	}
	if key == "FileID" {
		data, err = db.Query("SELECT * FROM orders WHERE FileID=?", value)
		if err != nil {
			panic(err)
		}
	} else if key == "OrderID" {
		data, err = db.Query("SELECT * FROM orders WHERE OrderID=?", value)
		if err != nil {
			panic(err)
		}
	} else if key == "Date" {
		data, err = db.Query("SELECT * FROM orders WHERE Date=?", value)
		if err != nil {
			panic(err)
		}
	}
	var ord Order
	for data.Next() {
		data.Scan(&ord.FileID, &ord.OrderID, &ord.Date, &ord.CustomerID, &ord.CustomerName, &ord.PizzaID, &ord.PizzaName, &ord.CardNum, &ord.Cost)
		c.HTML(http.StatusOK, "table.tmpl", gin.H{
			"FileID":       strconv.Itoa(ord.FileID),
			"OrderID":      ord.OrderID,
			"Date":         ord.Date,
			"CustomerID":   ord.CustomerID,
			"CustomerName": ord.CustomerName,
			"PizzaID":      ord.PizzaID,
			"PizzaName":    ord.PizzaName,
			"CardNum":      ord.CardNum,
			"Cost":         strconv.Itoa(ord.Cost),
		})
	}
}
