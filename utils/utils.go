package utils

import (
	"crypto/rand"
	"strconv"
)

func GenRanDigit() (string, error) {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}

func CheckDigit(cdn string) int {
	sum := 0
	for i := 0; i < len(cdn); i++ {
		d, _ := strconv.Atoi(string([]rune(cdn)[i]))
		if i%2 == 0 {
			sum = sum + (d * 1)
		} else {
			sum = sum + (d * 3)
		}
	}
	checkdigit := 0
	if sum%10 > 0 {
		checkdigit = 10 - (sum % 10)
	}
	return checkdigit
}
