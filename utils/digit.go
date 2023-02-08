package utils

import (
	"crypto/rand"
	"io"
	"strconv"
	"time"
)

// RandomDigit with custom length
func RandomDigit(length int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)

}

// GetCurrentYear get current year number
func GetCurrentYear() string {
	t := time.Now()
	return strconv.Itoa(t.Year())
}

// GetLastDigitYear return n last digit of year
func GetLastDigitYear(length int) string {
	code := GetCurrentYear()
	return code[len(code)-length:]
}
