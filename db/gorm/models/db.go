package models

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

type Gorm struct {
	db *gorm.DB
}

func (m *Gorm) DbInit() {
	dbName := os.Getenv("DATASOURCE_USERNAME")
	dbUrl := os.Getenv("DATASOURCE_URL")
	dbPass := os.Getenv("DATASOURCE_PASSWORD")
	dbUrl2 := strings.Replace(dbUrl, "jdbc:sqlserver://", "", 1)

	// "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	dsn := "sqlserver://" + dbName + ":" + dbPass + "@" + dbUrl2 + "?database=event"
	fmt.Println("db source name: ", dsn)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	m.db = db
	//return db
}

func (m *Gorm) DbClose() {
	db, err := m.db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	db.Close()
}

//type TokenInfoTable struct {
//	Name          string
//	Id            string
//	DefaultAmount string
//}
//
//func (m *Repository) Test() {
//	result := map[string]interface{}{}
//	m._db.Table("token_info_table").Take(&result)
//
//	fmt.Println("ok", result)
//}
