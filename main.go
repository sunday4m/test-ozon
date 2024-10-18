package main

import (
	"fmt"
	"test-ozon/service"
)

func main() {
	td := service.NewTestData()
	totalPhones := 100
	service.Generate(totalPhones, td)

	fmt.Printf("Сгенерировано %d номеров телефонов.\n", len(td.Phones))
	for phone := range td.Phones {
		fmt.Println(phone)
	}
}
