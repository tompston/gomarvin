package convert

import (
	"fmt"
	"strings"
)

func ConvertToTitle(i string) string {
	return strings.Title(i)
}
func ConvertToLowercase(i string) string {
	return strings.ToLower(i)
}
func ConvertToUppercase(i string) string {
	return strings.ToUpper(i)
}
func ConvertToLowercaseTitle(i string) string {
	return strings.Title(strings.ToLower(i))
}
func ConvertToCamelCase(i string) string {
	return ToCamel(i)
}
func ConvertToPlural(input string) string {
	return NewClient().Plural(input)
}
func ConvertToLowercasePlural(i string) string {
	return NewClient().Plural(strings.ToLower(i))
}

// Convert the last value of a string to be a new value
//
//	some_var := ConvertLastCharTo("AAAAAA", "BBB")
//	// some_var = AAAAABBB
func ConvertLastCharTo(x string, last_value string) string {
	new_str := strings.TrimRight(x, x[len(x)-1:])
	return fmt.Sprintf("%s%s", new_str, last_value)
}
