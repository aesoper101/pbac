package condition

import (
	"github.com/aesoper101/pbac/pdp/condition/consts"
	"github.com/aesoper101/pbac/pdp/types"
	"strings"
)

type StringCondition struct {
	baseCondition
	name string
	// compareFunc 比较函数, a为条件值, b为请求值
	compareFunc func(a, b interface{}) bool
}

func newStringCondition(name string, key string, values []interface{}, compareFunc func(a, b interface{}) bool) (types.Condition, error) {
	return &StringCondition{
		baseCondition: baseCondition{
			key:    key,
			values: values,
		},
		name:        name,
		compareFunc: compareFunc,
	}, nil
}

// newStringEqualsCondition StringEquals 精确匹配，区分大小写, 请求值与任意一个条件值相同（区分大小写）
func newStringEqualsCondition(key string, values []interface{}) (types.Condition, error) {
	return newStringCondition(consts.StringEquals, key, values, func(a, b interface{}) bool {
		return isString(a) && isString(b) && a == b
	})
}

// newStringNotEqualsCondition StringNotEquals 否定匹配,请求值与所有条件值都不同（区分大小写）
func newStringNotEqualsCondition(key string, values []interface{}) (types.Condition, error) {
	return newStringCondition(consts.StringNotEquals, key, values, func(a, b interface{}) bool {
		return isString(a) && isString(b) && a != b
	})
}

// newStringEqualsIgnoreCaseCondition StringEqualsIgnoreCase 请求值与任意一个条件值相同（不区分大小写）
func newStringEqualsIgnoreCaseCondition(key string, values []interface{}) (types.Condition, error) {
	return newStringCondition(consts.StringEqualsIgnoreCase, key, values, func(a, b interface{}) bool {
		return isString(a) && isString(b) && strings.EqualFold(a.(string), b.(string))
	})
}

// newStringNotEqualsIgnoreCaseCondition StringNotEqualsIgnoreCase 请求值与所有条件值都不同（不区分大小写）
func newStringNotEqualsIgnoreCaseCondition(key string, values []interface{}) (types.Condition, error) {
	return newStringCondition(consts.StringNotEqualsIgnoreCase, key, values, func(a, b interface{}) bool {
		return isString(a) && isString(b) && !strings.EqualFold(a.(string), b.(string))
	})
}

// newStringLikeCondition StringLike 请求值与任意一个条件值匹配（区分大小写）
func newStringLikeCondition(key string, values []interface{}) (types.Condition, error) {
	return newStringCondition(consts.StringLike, key, values, func(a, b interface{}) bool {
		return isString(a) && isString(b) && stringMatch(b.(string), a.(string))
	})
}

// newStringNotLikeCondition StringNotLike 请求值与所有条件值都不匹配（区分大小写）
func newStringNotLikeCondition(key string, values []interface{}) (types.Condition, error) {
	return newStringCondition(consts.StringNotLike, key, values, func(a, b interface{}) bool {
		return isString(a) && isString(b) && !stringMatch(b.(string), a.(string))
	})
}

func (c *StringCondition) GetName() string {
	return c.name
}

func (c *StringCondition) Evaluate(ctxValue interface{}, ctx types.EvalContextor) bool {
	return c.forOr(ctxValue, ctx, c.compareFunc)
}
