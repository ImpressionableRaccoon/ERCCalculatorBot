package decoder

import (
	"errors"
	"fmt"
	"math/bits"
	"strconv"
)

const (
	magicCode = 0x0E010A11
)

var (
	ErrInvalidNumberOfCharacters = errors.New("invalid number of characters")
	ErrWrongERCFormat            = errors.New("wrong ERC format")
)

func Decode(input string) (string, error) {
	if len(input) != 16 {
		return "", ErrInvalidNumberOfCharacters
	}

	f, err := strconv.ParseUint(input[:8], 16, 32)
	if err != nil {
		return "", ErrWrongERCFormat
	}
	first := uint32(f)

	s, err := strconv.ParseUint(input[8:], 16, 32)
	if err != nil {
		return "", ErrWrongERCFormat
	}
	second := uint32(s)

	res := first ^ bits.Reverse32(second) - magicCode

	return fmt.Sprintf("%08X", res), nil
}
