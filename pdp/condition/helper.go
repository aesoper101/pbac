package condition

import (
	"github.com/golang-module/carbon/v2"
	"github.com/shopspring/decimal"
	"regexp"
	"strconv"
	"strings"
)

// castNumber 将值转换为decimal.Decimal
func castNumber(fromValue string) (decimal.Decimal, bool) {
	var d decimal.Decimal

	d, err := decimal.NewFromString(fromValue)
	return d, err == nil
}

func stringMatch(a, b string) bool {
	pattern := "^" +
		strings.Replace(a, "[\\-\\[\\]]", "\\$&", -1) +
		strings.Replace(a, "\\*", ".*", -1) +
		strings.Replace(a, "\\?", ".", -1) +
		"$"

	matched, _ := regexp.MatchString(pattern, b)
	return matched
}

func isBool(value string) bool {
	v := strings.ToLower(value)
	return v == "true" || v == "false"
}

func isString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func castCarbon(value interface{}) (carbon.Carbon, bool) {
	switch value := value.(type) {
	case string:
		c := carbon.Parse(value)
		return c, c.IsValid()
	default:
		return carbon.NewCarbon(), false
	}
}

func getBoolString(value interface{}) (string, bool) {
	switch v := value.(type) {
	case bool:
		return strings.ToLower(strconv.FormatBool(v)), true
	case string:
		v1 := strings.ToLower(v)
		if v1 == "true" || v1 == "false" {
			return v1, true
		}

	}

	return "", false
}
