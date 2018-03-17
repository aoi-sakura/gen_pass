package main

import (
	"fmt"
	"flag"
	"strconv"

	cryptorand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

const (
	baseLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	symbols = ".-_@!#$%&()*+,/;<=>?[]^`{|}~"

	defaultLength = 16
)

var (
	lengthOpt = flag.Int("i", defaultLength, "password length.")
	noSymbolsOpt = flag.Bool("s", false, "generate password without all symbols.")
	specifiedSymbolsOpt = flag.String("S", "", "generate password with specified symbols.")
)

func getNumOfSquare(n int) uint {
	var i uint = 0
	for i = 0; n >= 1; i++ {
		n >>= 1
	}
	return i
}

func randSeed() int64 {
	var seed int64
	err := binary.Read(cryptorand.Reader, binary.LittleEndian, &seed)
	if err != nil {
		panic(err)
	}
	return seed
}

func generatePassword(n int, letters string) string {
	bit := getNumOfSquare(len(letters))
	var mask int64 = 1<<bit - 1
	var counterMax uint = 63 / bit
	var randomCounter uint = 0
	var randomVal int64

	b := make([]byte, n)
	for i := n - 1; i >= 0; {
		if randomCounter == 0 {
			randomVal = rand.Int63()
			randomCounter = counterMax
		}
		index := int(randomVal & mask)
		if index < len(letters) {
			b[i] = letters[index]
			i--
		}
		randomVal >>= bit
		randomCounter--
	}

	return string(b)
}

func main() {
	letters := baseLetters
	length := defaultLength

	flag.Parse()
	rand.Seed(randSeed())

	if flag.Arg(0) != "" {
		lengthValue, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			lengthValue = defaultLength
		}
		length = lengthValue
	} else {
		length = *lengthOpt
	}

	if *specifiedSymbolsOpt != "" {
		letters = letters + *specifiedSymbolsOpt
	} else if ! *noSymbolsOpt {
		letters = letters + symbols
	}

	fmt.Println(generatePassword(length, letters))
}
