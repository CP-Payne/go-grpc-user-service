package validation

import (
	"errors"
	"fmt"
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

// ValidateStringWithSpace ensures the string contains only alphabetic characters and spaces.
func ValidateStringWithSpace(value string) error {
	if value == "" {
		return nil
	}
	// Regular expression to match alphabetic characters and spaces
	if matched, _ := regexp.MatchString(`^[a-zA-Z\s]+$`, value); !matched {
		return errors.New("invalid string")
	}
	return nil
}

// ValidatePhone ensures the phone number contains only digits and specific characters (like +, -).
func ValidatePhone(value string) error {
	if value == "" {
		return nil
	}
	if matched, _ := regexp.MatchString(`^\+?[0-9\-]+$`, value); !matched {
		return errors.New("invalid phone number")
	}
	return nil
}

// ValidateHeight ensures the height is within a reasonable range.
func ValidateHeight(value float32) error {
	if value < 0 || value > 3 {
		return errors.New("invalid height")
	}
	return nil
}

// ValidateID ensures that the ID contains only integers greater than or equal to 1.
func ValidateID(id int32) error {
	if id < 1 {
		return errors.New("ID cannot be less than 1")
	}
	return nil
}

// ValidateIDs ensures that the list of IDs contains only integers greater than or equal to 1.
func ValidateIDs(ids []int32) error {
	if len(ids) == 0 {
		return errors.New("no IDs provided")
	}
	for _, id := range ids {
		if err := ValidateID(id); err != nil {
			return fmt.Errorf("invalid ID: %v", err)
		}
	}
	return nil
}
