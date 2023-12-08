package controller

import (
	"fmt"

	"github.com/SalhSThai/go-first-classes/src/repo"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	db := repo.Database()
	rows, err := db.Query("SELECT * FROM your_table")
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	fmt.Println("rows", rows)
	// for rows.Next() {
	// 	var column1, column2 string
	// 	err := rows.Scan(&column1, &column2)
	// 	if err != nil {
	// 		fmt.Println("err", err.Error())
	// 	}
	// 	fmt.Printf("Column1: %s, Column2: %s\n", column1, column2)
	// }

	c.JSON(200, c.Request.Body)
}
