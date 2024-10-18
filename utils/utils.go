package utils

import "math/rand"

func RandPhone() int {
	return 89000000000 + rand.Intn(999999999)
}
