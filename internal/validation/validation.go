package validation

import (
	"errors"
	"regexp"
)

// ValidateString ensures the string contains only alphabetic characters.
func ValidateString(value string) error {
	if value == "" {
		return nil
	}

	if matched, _ := regexp.MatchString(`^[a-zA-z]+$`, value); !matched {
		return errors.New("invalid string")
	}

	return nil
}
