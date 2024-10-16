package main

import (
	"fmt"
	"test-ozon/service"
)

func main() {
	td := service.NewTestData()
	td.GeneratePhones(100)
	fmt.Println(len(td.Phones))
}
