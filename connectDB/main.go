package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WebPage struct {
	URL      string `gorm:"column:url"`
	HTML     string `gorm:"column:html"`
	Title    string `gorm:"column:title"`
	Category string `gorm:"column:category"`
}

func (WebPage) TableName() string {
	return "webpage" // Set the table name explicitly
}

var db *gorm.DB

func main() {
	ConnectDB()
	p := WebPage{
		URL:      "test",
		Title:    "test",
		Category: "test",
	}
	p.insertDB()
}

func ConnectDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/scraper?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}

func (p *WebPage) insertDB() {
	result := db.Create(&p) // Insert the record into the database
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Record inserted successfully")
}
