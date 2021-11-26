package isbn

import "strings"

var (
	ISBN_VALID_LENGHTS = map[int]bool{
		10: true,
		13: true,
	}
)

type Validator struct {
}

func (v *Validator) Validate(isbnNumber string) bool {
	digits := strings.Split(isbnNumber, "")

	if !validateInput(digits) {
		return false
	}

	return true
}

func validateInput(digits []string) bool {
	if ok := ISBN_VALID_LENGHTS[len(digits)]; !ok {
		return false
	}

	firstThreedigits := strings.Join(digits[0:3], "")

	if len(digits) == 13 && firstThreedigit != "978" {
		return false
	}
	return true
}
