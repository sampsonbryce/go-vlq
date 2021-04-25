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
