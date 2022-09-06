package regexes

import (
	"fmt"
	"regexp"
)

func IsStandardFloat(value float64) bool {
	stringValue := fmt.Sprintf("%f", value)
	last3 := stringValue[len(stringValue)-3:]
	if last3 != "000" {
		return false
	}
	formattedValue := stringValue[0 : len(stringValue)-3]
	pattern := `^[0-9]{1,10}[.][0-9]{3}$`
	matched, err := regexp.MatchString(pattern, formattedValue)
	if err != nil {
		return false
	}
	return matched
}
