package vlq

import "testing"

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestDecode(t *testing.T) {
	toTest := make(map[string][]int)

	toTest["AAAA"] = []int{0, 0, 0, 0}
	toTest["EAAgB"] = []int{2, 0, 0, 16}
	toTest["mBAAD"] = []int{19, 0, 0, -1}
	toTest["SAAa"] = []int{9, 0, 0, 13}
	toTest["oE"] = []int{68}
	toTest["4mBAEA"] = []int{620, 0, 2, 0}
	toTest["D"] = []int{-1}

	for mapping, expected := range toTest {

		t.Run(mapping, func(t *testing.T) {
			result := Decode(mapping)

			if !Equal(result, expected) {
				t.Errorf("Mapping %s expected result %v but got %v", mapping, expected, result)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	toTest := make(map[string][]int)

	toTest["AAAA"] = []int{0, 0, 0, 0}
	toTest["EAAgB"] = []int{2, 0, 0, 16}
	toTest["mBAAD"] = []int{19, 0, 0, -1}
	toTest["SAAa"] = []int{9, 0, 0, 13}
	toTest["oE"] = []int{68}
	toTest["4mBAEA"] = []int{620, 0, 2, 0}
	toTest["D"] = []int{-1}

	for expected, input := range toTest {

		t.Run(expected, func(t *testing.T) {
			result := Encode(input)

			if expected != result {
				t.Errorf("Input %v expected result %v but got %v", input, expected, result)
			}
		})
	}
}
