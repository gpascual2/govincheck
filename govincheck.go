/*
MIT License

Copyright (c) 2017 Guillermo Pascual

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package govincheck

import (
	"math"
	"strconv"
	"strings"
)

// VinCheck - Vehicle Identification Number check digit validation
// Parameter: string - 17 digit VIN
// Return:
//    1- boolean - Validity flag. Set to true if VIN check digit is correct, false otherwise.
//    2- string - Valid VIN. Same VIN passed as parameter but with the correct check digit on it.
func VinCheck(vin string) (bool, string) {
	var valid = false
	vin = strings.ToUpper(vin)
	var retVin = vin

	if len(vin) == 17 {
		traSum := transcodeDigits(vin)
		checkNum := math.Mod(float64(traSum), 11)
		var checkDigit byte
		if checkNum == 10 {
			checkDigit = byte('X')
		} else {
			checkDigitTemp := strconv.Itoa(int(checkNum))
			checkDigit = checkDigitTemp[len(checkDigitTemp)-1]
		}
		if retVin[8] == checkDigit {
			valid = true
		}
		retVin = retVin[:8] + string(checkDigit) + retVin[9:]
	} else {
		valid = false
		retVin = ""
	}
	return valid, retVin
}

func transcodeDigits(vin string) int {
	var digitSum = 0
	var code int
	for i, chr := range vin {
		code = 0

		switch chr {
		case 'A', 'J', '1':
			code = 1
		case 'B', 'K', 'S', '2':
			code = 2
		case 'C', 'L', 'T', '3':
			code = 3
		case 'D', 'M', 'U', '4':
			code = 4
		case 'E', 'N', 'V', '5':
			code = 5
		case 'F', 'W', '6':
			code = 6
		case 'G', 'P', 'X', '7':
			code = 7
		case 'H', 'Y', '8':
			code = 8
		case 'R', 'Z', '9':
			code = 9
		case 'I', 'O', 'Q':
			code = 0
		}
		switch i + 1 {
		case 1, 11:
			digitSum += code * 8
		case 2, 12:
			digitSum += code * 7
		case 3, 13:
			digitSum += code * 6
		case 4, 14:
			digitSum += code * 5
		case 5, 15:
			digitSum += code * 4
		case 6, 16:
			digitSum += code * 3
		case 7, 17:
			digitSum += code * 2
		case 8:
			digitSum += code * 10
		case 9:
			digitSum += code * 0
		case 10:
			digitSum += code * 9
		}
	}
	return digitSum
}
