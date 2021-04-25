package vlq

import (
	"strings"
)

// Decode takes a base64 VLQ value and converts
// it into an array of values
// Heavily insipired by https://github.com/mozilla/source-map
// and https://github.com/Rich-Harris/vlq
func Decode(mapping string) []int {
	// binary: 100000
	VLQ_BASE := 1 << 5

	// binary: 011111
	VLQ_BASE_MASK := VLQ_BASE - 1

	// binary: 100000
	VLQ_CONTINUATION_MASK := VLQ_BASE

	// binary: 000001
	VLQ_SIGN_MASK := 1

	BASE64 := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	values := []int{}

	var value int = 0
	shift := 0

	for i := 0; i < len(mapping); i++ {
		digit := strings.Index(BASE64, string(mapping[i]))

		value |= (digit & int(VLQ_BASE_MASK)) << shift

		continues := digit & int(VLQ_CONTINUATION_MASK)

		if continues > 0 {
			shift += 5
		} else {
			sign := value & VLQ_SIGN_MASK

			// Remove sign bit
			value = value >> 1

			if sign > 0 {
				value = -value
			}

			values = append(values, value)

			// Reset
			value = 0
			shift = 0
		}
	}

	return values
}

// Encode takes an int array and converts
// it into base64 VLQ string
// Heavily insipired by https://github.com/mozilla/source-map
// and https://github.com/Rich-Harris/vl
func Encode(values []int) string {
	result := ""
	for _, value := range values {
		result += encodeInteger(value)
	}

	return result
}

func encodeInteger(value int) string {

	BASE64 := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	result := ""

	// Shift and set sign bit
	if value < 0 {
		value = (-value << 1) | 1
	} else {
		value = value << 1
	}

	for {
		// Save the first 5 bits
		field := value & 31
		value = value >> 5

		// Check if we are still encoding
		if value > 0 {
			field = field | 32 // Set continuation bit
		}

		result += string(BASE64[field])

		if value == 0 {
			break
		}
	}

	return result
}
