package ultilities

import (
	"errors"
	"unicode"
)

// Password validates plain password against the rules defined below.
//
// upp: at least one upper case letter.
// low: at least one lower case letter.
// num: at least one digit.
// sym: at least one special character.
// tot: at least eight characters long.
// No empty string or whitespace.
func ValidatePassword(value interface{}) error {
	pass := value.(string)
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return errors.New("password must contain at least one upper case letter, one lower case letter, one digit, one special character and at least eight characters long")
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return errors.New("password must contain at least one upper case letter, one lower case letter, one digit, one special character and at least eight characters long")
	}

	return nil
}
