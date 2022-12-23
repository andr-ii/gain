package random

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const c_len = len(charset)

var r *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func Str(length int) string {
	count := 0
	result := make([]byte, length)
	for count < length {
		result[count] = charset[r.Intn(c_len)]
		count++
	}

	return string(result)
}

func Num(min, max int) int {
	return r.Intn(max-min) + min
}
