package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Gorm struct {
	_db *gorm.DB
}

func (m *Gorm) DbInit() {
	/*
		dbName := os.Getenv("DATASOURCE_USERNAME")
		dbUrl := os.Getenv("DATASOURCE_URL")
		dbPass := os.Getenv("DATASOURCE_PASSWORD")
		dbUrl2 := strings.Replace(dbUrl, "jdbc:sqlserver://", "", 1)
	*/

	// env sqlserver dsn
	//dsn := "sqlserver://" + dbName + ":" + dbPass + "@" + dbUrl2 + "?database=event"
	// sqlserver dsn
	//dsn := "sqlserver://chain:block123!@172.16.28.98:1433?database=event"
	//db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	// marinaDB dsn
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("db source name: ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.AutoMigrate(&Token{})
	fmt.Println("db name: ", db)
	//db.Set("gorm:table_options", "")

	if err != nil {
		log.Fatalln(err)
	}

	m._db = db
}

func (m *Gorm) DbClose() {
	db, err := m._db.DB()
	//db, err := m.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}

	db.Close()
}

//func (m *Gorm) DbTokensFind(t *[]Token) {
//	m._db.Table("token_info_table").Find(t)
//	//m.DB.Table("token_info_table").Find(t)
//}
//
//func (m *Gorm) DbTokenFind(t *Token, id string) error {
//	if err := m._db.Table("token_info_table").Where("id = ?", id).First(t).Error; err != nil {
//		return err
//	}
//	return nil
//}
