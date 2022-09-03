package jobdata

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func Jobdata() {
	fName := "jobdata/data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file,error:%q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})
	for i := 0; i < 5; i++ {
		fmt.Printf("Scrapping Page: %d\n", i)
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}
	log.Printf("Scraping Completed\n")
	log.Println(c)

}
