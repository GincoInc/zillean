package zillean

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

var unitMap = map[string]string{
	"qa":  "1",
	"li":  "1000000",
	"zil": "1000000000000",
}

// FromQa converts a number in Qa to in specified unit.
func FromQa(qa string, unit string, pad bool) (string, error) {
	if unit == "qa" {
		return qa, nil
	}
	if base, ok := unitMap[unit]; ok {
		baseNumDecimals := len(base) - 1
		_qa := strToBigInt(qa)
		_base := strToBigInt(base)

		_fraction := &big.Int{}
		fraction := _fraction.Mod(_fraction.Abs(_qa), _base).String()
		for len(fraction) < baseNumDecimals {
			fraction = "0" + fraction
		}

		if !pad {
			r := regexp.MustCompile("^([0-9]*[1-9]|0)(0*)")
			fraction = r.FindAllStringSubmatch(fraction, -1)[0][1]
		}

		_whole := &big.Int{}
		whole := _whole.Quo(_qa, _base).String()

		if fraction == "0" {
			return whole, nil
		}

		return whole + "." + fraction, nil
	}

	return "", errors.New("the specified unit is invalid")
}

// ToQa converts a number in specified unit to in Qa.
func ToQa(num string, unit string) (string, error) {
	if base, ok := unitMap[unit]; ok {
		baseNumDecimals := len(base) - 1
		_base := strToBigInt(base)

		inputStr := num
		isNegative := string(num[0]) == "-"
		if isNegative {
			inputStr = num[1:]
		}

		if inputStr == "." {
			return "", errors.New("cannot convert to Qa")
		}

		var whole, fraction string
		comps := strings.Split(inputStr, ".")
		fmt.Println(len(comps))
		switch len(comps) {
		case 1:
			whole = comps[0]
			fraction = "0"
		case 2:
			whole = comps[0]
			fraction = comps[1]
		default:
			return "", errors.New("cannot convert to Qa")
		}

		if whole == "" {
			whole = "0"
		}

		if fraction == "" {
			fraction = "0"
		}

		if len(fraction) > baseNumDecimals {
			return "", errors.New("cannot convert to Qa")
		}

		for len(fraction) < baseNumDecimals {
			fraction += "0"
		}

		_whole := strToBigInt(whole)
		_fraction := strToBigInt(fraction)

		wei := &big.Int{}
		wei.Add(wei.Mul(_whole, _base), _fraction)

		if isNegative {
			wei.Neg(wei)
		}

		return wei.String(), nil

	}
	return "", errors.New("the specified unit is invalid")
}
