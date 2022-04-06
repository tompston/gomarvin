package convert

import (
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
