package luhn

import (
	"errors"
	"fmt"
	"os"
)

var ERR_INVALID_LENGTH_SHORT = errors.New("CC Number is too short < 16 characters")
var ERR_INVALID_LENGTH_LONG = errors.New("CC Number is too long > 19 characters")
var ERR_INVALID_NUMBER = errors.New("CC Number contains unexpected non-numeric characters")
var ERR_INVALID_LUHN = errors.New("CC Number has an invalid LUHN 10")

const (
	MIN_LENGTH = 16
	MAX_LENGTH = 19
)

var Debug bool = false

func debug(format string, args ...interface{}) {
	if !Debug {
		return
	}
	fmt.Fprintf(os.Stderr, "[DEBUG]: "+format, args)
}

var oddLookup = [10]int{0,2,4,6,8,1,3,5,7,9}

func ValidLuhn(cardnumber string) error {
	debug("new number %s\n", cardnumber)
	cardlen := len(cardnumber)
	if cardlen < MIN_LENGTH {
		return ERR_INVALID_LENGTH_SHORT
	}
	if cardlen > MAX_LENGTH {
		return ERR_INVALID_LENGTH_LONG
	}
	sum := 0
	digits := []byte(cardnumber)
	for i, iMax := 0, len(digits); i < iMax; i += 1 {
		working := int(digits[i] - 48)
		if working < 0 || working > 9 {
			return ERR_INVALID_NUMBER
		}
		debug("digit was %d\n", working)
		if i%2 == 0 {
			working = oddLookup[working]
		}
		debug("Working was %d\n", working)
		sum += working
	}
	if sum%10 == 0 {
		return nil
	}
	debug("LUHN was %d\n", sum%10)
	return ERR_INVALID_LUHN
}
