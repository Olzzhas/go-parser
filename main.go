package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	var data [][]string

	c.OnHTML(".row", func(e *colly.HTMLElement) {
		e.ForEach(".row__top", func(_ int, rowTop *colly.HTMLElement) {
			rank := rowTop.ChildText(".row-cell.rank span")
			instagram := rowTop.ChildText(".row-cell.contributor .contributor__name-content")
			name := rowTop.ChildText(".row-cell.contributor .contributor__title")
			followers := rowTop.ChildText(".row-cell.subscribers")
			country := rowTop.ChildText(".row-cell.audience")
			authentic := rowTop.ChildText(".row-cell.authentic")
			engagement := rowTop.ChildText(".row-cell.engagement")

			rowData := []string{rank, instagram, name, followers, country, authentic, engagement}
			data = append(data, rowData)
		})
	})

	// Посещаем стартовую страницу
	err := c.Visit("https://hypeauditor.com/top-instagram-all-russia/")
	if err != nil {
		log.Fatal(err)
	}

	// Создаем CSV файл
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Создаем записыватель CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Записываем данные в CSV
	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Данные успешно сохранены в output.csv")
}
