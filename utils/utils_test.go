package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandPhoneRange(t *testing.T) {
	phone := RandPhone()

	assert.GreaterOrEqual(t, phone, 89000000000, "Ожидалось, что номер телефона будет больше или равен 89000000000")
	assert.LessOrEqual(t, phone, 89999999999, "Ожидалось, что номер телефона будет меньше или равен 89999999999")
}

func TestRandPhoneMultipleCalls(t *testing.T) {
	phone1 := RandPhone()
	phone2 := RandPhone()

	// Проверяем, что два последовательно сгенерированных номера не совпадают (не гарантирует случайность, но помогает удостовериться, что функция не возвращает одно и то же значение)
	assert.NotEqual(t, phone1, phone2, "Ожидалось, что два последовательно сгенерированных номера будут различными")
}
