package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
)

var (
	lowerChars   = "abcdedfghijklmnopqrst"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChars = "!@#$%&*"
	numbers      = "0123456789"
	allChars     = lowerChars + upperChars + specialChars + numbers
)

func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder
	for i := 0; i < minSpecialChar; i++ {
		random, _ := crand.Int(crand.Reader, big.NewInt(int64(len(specialChars))))
		password.WriteString(string(specialChars[random.Int64()]))
	}
	for i := 0; i < minNum; i++ {
		random, _ := crand.Int(crand.Reader, big.NewInt(int64(len(numbers))))
		password.WriteString(string(numbers[random.Int64()]))
	}
	for i := 0; i < minUpperCase; i++ {
		random, _ := crand.Int(crand.Reader, big.NewInt(int64(len(upperChars))))
		password.WriteString(string(upperChars[random.Int64()]))
	}
	numLowerCase := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < numLowerCase; i++ {
		random, _ := crand.Int(crand.Reader, big.NewInt(int64(len(allChars))))
		password.WriteString(string(allChars[random.Int64()]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
func main() {
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 20
	m := make(map[string]int)
	for i := 0; i < 100000; i++ {
		password := GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
		m[password]++
	}
	for k, v := range m {
		if v > 1 {
			fmt.Println(k)
		}
	}
}
