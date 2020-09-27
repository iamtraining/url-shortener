package libraries

import (
	"math/rand"
)

func Generate(n int) string {
	const alpha = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(bytes)
}
