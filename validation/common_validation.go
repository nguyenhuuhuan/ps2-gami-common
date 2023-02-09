package validation

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"gitlab.id.vin/gami/gami-common/models"
)

const (
	excludeSpecialCharacters               = "^[^~!@#$%^&*()+=\\/*+?,`<>%|:;\"'{}$]*$"
	excludeSpecialCharactersAndWhiteSpaces = "^[^~!@#$% ^&*()+=\\/*+?,`<>%|:;\"'{}$]*$"
	excludeSpecialCharactersAndUnderline   = "^[^~!@#$%_^&*()+=\\/*+?,`<>%|:;\"'{}$]*$"
)

var excludeSpecialChars = regexp.MustCompile(excludeSpecialCharacters)
var excludeSpecialCharsAndWhiteSpaces = regexp.MustCompile(excludeSpecialCharactersAndWhiteSpaces)
var excludeSpecialCharsAndUnderline = regexp.MustCompile(excludeSpecialCharactersAndUnderline)

// Match will check the value of fl matching with regexp2 or not.
func Match(reg *regexp.Regexp, fl validator.FieldLevel) bool {
	return reg.MatchString(fl.Field().String())
}

// CustomDateTimeFormatValidation Validate if string is match with date time format ISO 8601
func CustomDateTimeFormatValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse(models.DateFormat, fl.Field().String())
	return err == nil
}

// ExcludeSpecialChars will accept a string not containing any special characters.
func ExcludeSpecialChars(fl validator.FieldLevel) bool {
	return Match(excludeSpecialChars, fl)
}

// ExcludeSpecialCharsAndWhiteSpaces will accept a string not containing any special characters and white spaces.
func ExcludeSpecialCharsAndWhiteSpaces(fl validator.FieldLevel) bool {
	return Match(excludeSpecialCharsAndWhiteSpaces, fl)
}

func ExcludeSpecialCharsAndUnderline(fl validator.FieldLevel) bool {
	return Match(excludeSpecialCharsAndUnderline, fl)
}
