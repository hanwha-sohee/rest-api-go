package models

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

type Gorm struct {
	_db *gorm.DB
	//DB *gorm.DB
}

func (m *Gorm) DbInit() {
	//dbName := os.Getenv("DATASOURCE_USERNAME")
	//dbUrl := os.Getenv("DATASOURCE_URL")
	//dbPass := os.Getenv("DATASOURCE_PASSWORD")
	//dbUrl2 := strings.Replace(dbUrl, "jdbc:sqlserver://", "", 1)

	// "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	//dsn := "sqlserver://" + dbName + ":" + dbPass + "@" + dbUrl2 + "?database=event"
	dsn := "sqlserver://chain:block123!@172.16.28.98:1433?database=event"
	fmt.Println("db source name: ", dsn)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Token{})
	//fmt.Println("db source name: ", db)
	//db.Set("gorm:table_options", "")

	if err != nil {
		log.Fatalln(err)
	}

	m._db = db
	//m.DB = db
	//return db
}

func (m *Gorm) DbClose() {
	db, err := m._db.DB()
	//db, err := m.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}

	db.Close()
}

func (m *Gorm) DbTokenFind(t *[]Token) {
	m._db.Table("token_info_table").Find(t)
	//m.DB.Table("token_info_table").Find(t)
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
