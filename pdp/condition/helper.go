package condition

import (
	"github.com/golang-module/carbon/v2"
	"github.com/shopspring/decimal"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// castNumber 将值转换为decimal.Decimal
func castNumber(fromValue interface{}) (decimal.Decimal, bool) {
	var d decimal.Decimal

	switch value := fromValue.(type) {
	case int, int8, int16, int32, int64:
		d = decimal.NewFromInt(value.(int64))
	case uint, uint8, uint16, uint32, uint64:
		d = decimal.NewFromInt(int64(value.(uint64)))
	case float32, float64:
		d = decimal.NewFromFloat(value.(float64))
	case string:
		var err error
		d, err = decimal.NewFromString(value)
		if err != nil {
			return decimal.Zero, false
		}
	default:
		return decimal.Zero, false
	}

	return d, true
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
	case time.Time:
		c := carbon.CreateFromStdTime(value)
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
