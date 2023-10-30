// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/gocolly/colly"
// )

// func main() {
// 	c := colly.NewCollector()

// 	var title, html string

// 	// Crawl title
// 	c.OnHTML("title", func(e *colly.HTMLElement) {
// 		title = e.Text
// 		fmt.Println("Title:", title)
// 	})

// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		e.Request.Visit(e.Attr(("href")))
// 	})

// 	// Crawl full HTML
// 	c.OnResponse(func(r *colly.Response) {
// 		html = string(r.Body)
// 		fmt.Println("Full HTML:", html)
// 	})

// 	// Set limit rules
// 	c.Limit(&colly.LimitRule{
// 		DomainGlob:  "*",
// 		Parallelism: 2,               // Số request được phép cùng lúc
// 		RandomDelay: 2 * time.Second, // Thời gian delay giữa các request
// 	})

// 	// Visiting the URL
// 	err := c.Visit("https://umcclinic.com.vn/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Ask user for file path
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter file path to save data (including filename.txt): ")
// 	filePath, _ := reader.ReadString('\n')
// 	filePath = strings.TrimSpace(filePath)

// 	// Create and write to the file
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)
// 	writer.WriteString("Title: " + title + "\n\n")
// 	writer.WriteString("Full HTML: " + html)
// 	writer.Flush()

// 	fmt.Println("Data saved successfully.")
// }
