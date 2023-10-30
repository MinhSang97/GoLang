package main

import (
	"log"
	"regexp"
	"strings"
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
				Category: strings.Join(paragraphs[i], "\n"),
			}
			insertDB(p)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

func ConnectDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/scraper?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}

func CrawlContent(url string) (string, []string) {
	c := colly.NewCollector()

	var title string
	var paragraphs []string

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML("span, strong, div, a, b, h1, h2, h3, h4, h5, h6, p", func(e *colly.HTMLElement) {
		text := e.Text
		if text != "" {
			// Use regular expressions to clean up spaces
			text = regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(text), " ")
			paragraphs = append(paragraphs, text)
		}
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 4 * time.Second,
	})

	err := c.Visit(url)
	if err != nil {
		log.Println("Error visiting URL:", url)
	}

	return title, paragraphs
}

func Crawler() ([]string, []string, [][]string) {
	c := colly.NewCollector()

	var collectedURLs []string
	var collectedTitles []string
	var collectedParagraphs [][]string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		collectedURLs = append(collectedURLs, link)

		title, paragraphs := CrawlContent(link)
		collectedTitles = append(collectedTitles, title)
		collectedParagraphs = append(collectedParagraphs, paragraphs)
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 4 * time.Second,
	})

	err := c.Visit("https://ump.edu.vn/")
	if err != nil {
		log.Fatal(err)
	}

	minLength := min(len(collectedURLs), len(collectedTitles), len(collectedParagraphs))
	collectedURLs = collectedURLs[:minLength]
	collectedTitles = collectedTitles[:minLength]
	collectedParagraphs = collectedParagraphs[:minLength]

	return collectedURLs, collectedTitles, collectedParagraphs
}

func insertDB(p WebPage) {
	result := db.Exec("INSERT INTO webpage (url, html, title, category) VALUES (?, ?, ?, ?)", p.URL, p.HTML, p.Title, p.Category)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Record inserted successfully")
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
