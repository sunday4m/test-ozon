package service

import (
	"sync"
	"test-ozon/utils"
)

type TestData struct {
	Phones map[int]struct{}
	Mu     sync.Mutex // для синхронизации
}

func NewTestData() *TestData {
	return &TestData{
		Phones: make(map[int]struct{}),
	}
}

func (td *TestData) Add(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		phone := utils.RandPhone()

		td.Mu.Lock()

		// Проверим уникальность номера
		if _, exists := td.Phones[phone]; !exists {
			td.Phones[phone] = struct{}{}
			td.Mu.Unlock()
			return
		}

		td.Mu.Unlock()
	}
}

func Generate(n int, td *TestData) {
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go td.Add(&wg)
	}

	wg.Wait()
}
