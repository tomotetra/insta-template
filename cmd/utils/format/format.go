package format

import (
	"math/big"
	"strings"
)

// RemoveQuotes removes surrounding double-quotes
func RemoveQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

// RatToString converts rational numbers to string
// preserve rational if numerator or denometer equals to 1.
// convert to float string otherwise.
func RatToString(r *big.Rat) string {
	var stringVal string
	if r.Num().Cmp(big.NewInt(1))*r.Denom().Cmp(big.NewInt(1)) == 0 {
		stringVal = r.RatString()
	} else {
		stringVal = r.FloatString(1)
	}
	return stringVal
}

// DateTimeString formats datetime string
// i.e. 2019:03:10 20:00:00 -> 2019.03.10
func DateTimeString(d string) string {
	return strings.Replace(strings.Split(d, " ")[0], ":", ".", -1)
}
