package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gocolly/colly"
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
	urls, titles, paragraphs := Crawler()
	ConnectDB()

	var wg sync.WaitGroup

	for i := range urls {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			p := WebPage{
				URL:      urls[i],
				Title:    titles[i],
				Category: paragraphs[i],
			}
			p.insertDB()
		}(i)
	}

	// Chờ tất cả goroutine hoàn thành
	wg.Wait()
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

func Crawler() ([]string, []string, []string) {
	c := colly.NewCollector()

	var title, html string
	var collectedURLs []string
	var collectedTitles []string
	var collectedParagraphs []string

	// Crawl title
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
		collectedTitles = append(collectedTitles, title) // Thêm title vào slice
		fmt.Println("Title:", title)

		c.OnHTML(fmt.Sprintf("p:contains('%s')", title), func(e *colly.HTMLElement) {
			paragraph := e.Text
			collectedParagraphs = append(collectedParagraphs, paragraph) // Thêm paragraph vào slice
			fmt.Println("Paragraph:", paragraph)
		})
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		collectedURLs = append(collectedURLs, link) // Thêm link vào slice
		e.Request.Visit(link)
		fmt.Println("url:", link)
	})

	for range collectedTitles {
		c.OnHTML("p", func(e *colly.HTMLElement) {
			paragraph := e.Text
			collectedParagraphs = append(collectedParagraphs, paragraph)
			fmt.Println("Paragraph:", paragraph)
		})
	}

	c.OnHTML("a"+"span"+"b", func(e *colly.HTMLElement) {
		paragraph := e.Text
		collectedParagraphs = append(collectedParagraphs, paragraph) // Thêm paragraph vào slice
		fmt.Println("Paragraph:", paragraph)
	})

	// Crawl full HTML
	c.OnResponse(func(r *colly.Response) {
		html = string(r.Body)
		fmt.Println("Full HTML:", html)
	})

	// Set limit rules
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,               // Số request được phép cùng lúc
		RandomDelay: 4 * time.Second, // Thời gian delay giữa các request
	})

	// Visiting the URL
	err := c.Visit("http://14.225.192.61/")
	if err != nil {
		log.Fatal(err)
	}
	// Đảm bảo cả ba slice có cùng độ dài
	minLength := min(len(collectedURLs), len(collectedTitles), len(collectedParagraphs))
	collectedURLs = collectedURLs[:minLength]
	collectedTitles = collectedTitles[:minLength]
	collectedParagraphs = collectedParagraphs[:minLength]

	return collectedURLs, collectedTitles, collectedParagraphs
}

func min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func (p *WebPage) insertDB() {
	result := db.Create(&p) // Insert the record into the database
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Record inserted successfully")
}
