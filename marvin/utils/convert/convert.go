package convert

import (
	"fmt"
	"strings"
)

func ToTitle(i string) string {
	return strings.Title(i)
}

func ToLowercase(i string) string {
	return strings.ToLower(i)
}

func ToUppercase(i string) string {
	return strings.ToUpper(i)
}

func ToLowercaseTitle(i string) string {
	return strings.Title(strings.ToLower(i))
}

func ToCamelCase(i string) string {
	return ToCamel(i)
}

func ToPlural(input string) string {
	return NewClient().Plural(input)
}

func ToLowercasePlural(i string) string {
	return NewClient().Plural(strings.ToLower(i))
}

func LastCharTo(x string, last_value string) string {
	new_str := strings.TrimRight(x, x[len(x)-1:])
	return fmt.Sprintf("%s%s", new_str, last_value)
}

func WrapInCurlyBraces(x string) string {
	return fmt.Sprintf("{" + x + "}")
}

func WrapInCurlyBracesWithAppendedString(x string, appended_string string) string {
	return fmt.Sprintf("{" + x + appended_string + "}")
}
