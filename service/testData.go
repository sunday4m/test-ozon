package service

import "math/rand"

type TestData struct {
	Phones map[int]struct{}
}

func NewTestData() *TestData {
	return &TestData{
		Phones: make(map[int]struct{}),
	}
}

func (td *TestData) GeneratePhones(n int) {
	for i := 0; i < n; i++ {
		for {
			phone := randPhone()
			// Проверка на уникальность номера
			if _, exists := td.Phones[phone]; !exists {
				td.Phones[phone] = struct{}{}
				break // Завершаем внутренний цикл, когда номер добавлен
			}
		}
	}
}

func randPhone() int {
	return 89090000000 + rand.Intn(9999999)
}