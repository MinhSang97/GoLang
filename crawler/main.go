package main

import (
	"app/crawler/dbutil"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// // Tạo một instance của Colly
	// c := colly.NewCollector()

	// // Khai báo slice để lưu trữ danh sách repositories
	// var repositories []string

	// // Cài đặt hàm callback cho việc tìm thấy một thẻ "a" có class là "text-bold"
	// c.OnHTML("a.text-bold", func(e *colly.HTMLElement) {
	// 	repositoryName := strings.TrimSpace(e.Text)
	// 	repositories = append(repositories, repositoryName)
	// })

	// // Cài đặt hàm callback cho việc tìm thấy một thẻ "a" có attribute href bắt đầu bằng "/username/"
	// c.OnHTML("a[href^='/']", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	if strings.HasPrefix(link, "/") && strings.Count(link, "/") == 2 {
	// 		// Thêm đường link đầy đủ vào mỗi repository
	// 		fullLink := "https://github.com" + link
	// 		fmt.Println("Repository Link:", fullLink)
	// 	}
	// })

	// // Cài đặt hàm callback cho việc in ra thông báo khi kết thúc crawl
	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("Crawl finished. Found repositories:")
	// 	for _, repo := range repositories {
	// 		fmt.Println(repo)
	// 	}
	// })

	// // Bắt đầu crawl từ trang chủ của GitHub
	// err := c.Visit("https://github.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	db := dbutil.ConnectDB()
	fmt.Println("Connected: ", db)
	crawler()
}

func crawler() {
	c := colly.NewCollector()

	// Crawl title
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Println("Title:", title)
	})

	// Crawl URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Crawl full HTML
	c.OnResponse(func(r *colly.Response) {
		html := string(r.Body)
		fmt.Println("Full HTML:", html)
	})

	// Visiting the URL
	err := c.Visit("https://ump.edu.vn/")
	if err != nil {
		log.Fatal(err)
	}
}
