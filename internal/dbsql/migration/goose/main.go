package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"task-management-api/internal/config"
	"task-management-api/internal/dbsql"

	"github.com/pressly/goose/v3"
)

func main() {
	fileMigratePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fileMigratePath = path.Join(fileMigratePath, "db")

	if _, err := os.Stat(fileMigratePath); os.IsNotExist(err) {
		fmt.Println(fileMigratePath, "does not exists.")
		return
	}

	config.InitConfig()

	err = dbsql.InitDB()
	if err != nil {
		panic(err)
	}

	dbConn := dbsql.GetDB()
	dbConnSql, err := dbConn.DB()
	if err != nil {
		log.Println(err)
	}

	db := dbConnSql

	// Migrate the database to the latest version
	if err = goose.SetDialect("mysql"); err != nil {
		log.Println("goose Upto:", err)
	}
	goose.SetVerbose(true)

	if !(strings.Contains(os.Getenv("DB_NAME"), "task_management_api")) {
		fmt.Println("Currently on testing, DB name must contain \"test\".")
		return
	}
	fmt.Println("Migrating DB:", os.Getenv("DB_NAME"))
	fmt.Printf("Confirm? [Y/n]: ")
	reader := bufio.NewReader(os.Stdin)
	res, err := reader.ReadString('\n')
	if err != nil {
		//log.Fatal(err)
		return
	}
	if strings.TrimSpace(res) != "Y" {
		return
	}

	//goose.Up(db, fileMigratePath)

	if err = goose.UpTo(db, fileMigratePath, 1); err != nil {
		log.Println("goose UpTo:", err)
	}

	// if err = goose.Down(db, fileMigratePath); err != nil {
	// 	log.Println("goose DownTo:", err)
	// }

	if err = goose.Status(db, fileMigratePath); err != nil {
		log.Println("goose Status:", err)
	}

	// Note: cannot seed purchasing_companies and rfq_categories since it's related to company
}
