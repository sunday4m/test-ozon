package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestDataTable(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{name: "Пустая TestData", expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := NewTestData()
			assert.NotNil(t, td, "Ожидалось, что TestData будет инициализирована, но получено nil")
			assert.NotNil(t, td.Phones, "Ожидалось, что карта Phones будет инициализирована, но получено nil")
			assert.Equal(t, tt.expected, len(td.Phones), "Ожидалось, что телефонов не будет")
		})
	}
}

func TestGeneratePhones(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected int
	}{
		{name: "Генерация 0 телефонов", num: 0, expected: 0},
		{name: "Генерация 1 телефона", num: 1, expected: 1},
		{name: "Генерация 10 телефонов", num: 10, expected: 10},
		{name: "Генерация 100 телефонов", num: 100, expected: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := NewTestData()
			td.GeneratePhones(tt.num)
			assert.Equal(t, tt.expected, len(td.Phones), "Неожиданное количество сгенерированных телефонов")
		})
	}
}

func TestUniquePhones(t *testing.T) {
	tests := []struct {
		name string
		num  int
	}{
		{name: "Генерация 1 телефона", num: 1},
		{name: "Генерация 10 телефонов", num: 10},
		{name: "Генерация 1000 телефонов", num: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := NewTestData()
			td.GeneratePhones(tt.num)

			seen := make(map[int]bool)
			for phone := range td.Phones {
				assert.False(t, seen[phone], "Обнаружен дублирующийся телефон: %d", phone)
				seen[phone] = true
			}
		})
	}
}

func TestValidPhoneNumbers(t *testing.T) {
	tests := []struct {
		name string
		num  int
	}{
		{name: "Генерация 0 телефонов", num: 0},
		{name: "Генерация 10 телефонов", num: 10},
		{name: "Генерация 1000 телефонов", num: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := NewTestData()
			td.GeneratePhones(tt.num)

			for phone := range td.Phones {
				assert.GreaterOrEqual(t, phone, 89090000000, "Телефон меньше допустимого диапазона")
				assert.LessOrEqual(t, phone, 89099999999, "Телефон больше допустимого диапазона")
			}
		})
	}
}

func TestGeneratePhonesMultipleCalls(t *testing.T) {
	tests := []struct {
		name      string
		call1     int
		call2     int
		expected  int
	}{
		{name: "Первый вызов 0, второй вызов 50", call1: 0, call2: 50, expected: 50},
		{name: "Первый вызов 50, второй вызов 50", call1: 50, call2: 50, expected: 100},
		{name: "Первый вызов 100, второй вызов 100", call1: 100, call2: 100, expected: 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := NewTestData()
			td.GeneratePhones(tt.call1)
			td.GeneratePhones(tt.call2)

			assert.Equal(t, tt.expected, len(td.Phones), "Неожиданное количество уникальных телефонов после нескольких вызовов")

			seen := make(map[int]bool)
			for phone := range td.Phones {
				assert.False(t, seen[phone], "Обнаружен дублирующийся телефон: %d", phone)
				seen[phone] = true
			}
		})
	}
}

func TestRandPhone(t *testing.T) {
	for i := 0; i < 100; i++ {
		phone := randPhone()
		assert.GreaterOrEqual(t, phone, 89090000000, "Телефон меньше допустимого диапазона")
		assert.LessOrEqual(t, phone, 89099999999, "Телефон больше допустимого диапазона")
	}
}