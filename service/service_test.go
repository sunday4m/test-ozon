package service

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUniquePhones(t *testing.T) {
	tests := []struct {
		name        string
		totalPhones int
	}{
		{
			name:        "Генерация 1 уникального номера",
			totalPhones: 1,
		},
		{
			name:        "Генерация 100 уникальных номеров",
			totalPhones: 100,
		},
		{
			name:        "Генерация 0 номеров",
			totalPhones: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := NewTestData()

			Generate(tt.totalPhones, td)

			assert.Equal(t, tt.totalPhones, len(td.Phones), "Ожидалось %d уникальных номеров, но получено %d", tt.totalPhones, len(td.Phones))

			seen := make(map[int]struct{})
			for phone := range td.Phones {
				_, exists := seen[phone]
				assert.False(t, exists, "Номер %d повторяется", phone)
				seen[phone] = struct{}{}
			}
		})
	}
}

// вообще я не уверен, что нужен этот тест, но я подумал, раз работаем с мьютексами, то можно чет подобное придумать
func TestParallelPhoneGeneration(t *testing.T) {
	tests := []struct {
		name        string
		totalPhones int
		threads     int
	}{
		{
			name:        "Параллельная генерация с 2 потоками",
			totalPhones: 100,
			threads:     2,
		},
		{
			name:        "Параллельная генерация с 20 потоками",
			totalPhones: 2000,
			threads:     20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := NewTestData()
			var wg sync.WaitGroup

			for i := 0; i < tt.threads; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					Generate(tt.totalPhones/tt.threads, td)
				}()
			}

			wg.Wait()

			assert.Equal(t, tt.totalPhones, len(td.Phones), "Ожидалось %d уникальных номеров, но получено %d", tt.totalPhones, len(td.Phones))

			seen := make(map[int]struct{})
			for phone := range td.Phones {
				_, exists := seen[phone]
				assert.False(t, exists, "Номер %d повторяется", phone)
				seen[phone] = struct{}{}
			}
		})
	}
}

func TestNewTestData(t *testing.T) {
	td := NewTestData()
	assert.NotNil(t, td, "Ожидалось, что TestData не будет nil")
	assert.Equal(t, 0, len(td.Phones), "Ожидалась пустая карта телефонов, но обнаружено %d элементов", len(td.Phones))
}
