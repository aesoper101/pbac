package condition

import (
	"github.com/aesoper101/pbac/pdp/condition/consts"
	"github.com/aesoper101/pbac/pdp/types"
)

type dateCondition struct {
	baseKeyedCondition
	name string
	// compareFunc 比较函数, a为条件值, b为请求值
	compareFunc func(a, b interface{}) bool
}

func newDateCondition(name string, key string, values []interface{}, compareFunc func(a, b interface{}) bool) KeyedCondition {
	return &dateCondition{
		baseKeyedCondition: baseKeyedCondition{
			key:    key,
			values: values,
		},
		name:        name,
		compareFunc: compareFunc,
	}
}

// newDateEqualsCondition DateEquals 请求值等于任意一个条件值。
func newDateEqualsCondition(key string, values []interface{}) KeyedCondition {
	return newDateCondition(consts.DateEquals, key, values, func(a, b interface{}) bool {
		aDate, aOK := castCarbon(a)
		bDate, bOK := castCarbon(b)
		return aOK && bOK && bDate.Eq(aDate)
	})
}

// newDateNotEqualsCondition DateNotEquals 请求值不等于所有条件值。
func newDateNotEqualsCondition(key string, values []interface{}) KeyedCondition {
	return newDateCondition(consts.DateNotEquals, key, values, func(a, b interface{}) bool {
		aDate, aOK := castCarbon(a)
		bDate, bOK := castCarbon(b)
		return aOK && bOK && !bDate.Eq(aDate)
	})
}

// newDateLessThanCondition DateLessThan 请求值小于所有条件值。
func newDateLessThanCondition(key string, values []interface{}) KeyedCondition {
	return newDateCondition(consts.DateLessThan, key, values, func(a, b interface{}) bool {
		aDate, aOK := castCarbon(a)
		bDate, bOK := castCarbon(b)
		return aOK && bOK && bDate.Lt(aDate)
	})
}

// newDateLessThanEqualsCondition DateLessThanEquals 请求值小于等于所有条件值。
func newDateLessThanEqualsCondition(key string, values []interface{}) KeyedCondition {
	return newDateCondition(consts.DateLessThanEquals, key, values, func(a, b interface{}) bool {
		aDate, aOK := castCarbon(a)
		bDate, bOK := castCarbon(b)
		return aOK && bOK && bDate.Lte(aDate)
	})
}

// newDateGreaterThanCondition DateGreaterThan 请求值大于所有条件值。
func newDateGreaterThanCondition(key string, values []interface{}) KeyedCondition {
	return newDateCondition(consts.DateGreaterThan, key, values, func(a, b interface{}) bool {
		aDate, aOK := castCarbon(a)
		bDate, bOK := castCarbon(b)
		return aOK && bOK && bDate.Gt(aDate)
	})
}

// newDateGreaterThanEqualsCondition DateGreaterThanEquals 请求值大于等于所有条件值。
func newDateGreaterThanEqualsCondition(key string, values []interface{}) KeyedCondition {
	return newDateCondition(consts.DateGreaterThanEquals, key, values, func(a, b interface{}) bool {
		aDate, aOK := castCarbon(a)
		bDate, bOK := castCarbon(b)
		return aOK && bOK && bDate.Gte(aDate)
	})
}

func (c *dateCondition) GetName() string {
	return c.name
}

func (c *dateCondition) Evaluate(ctxValue interface{}, requestCtx types.EvalContextor) bool {
	values := c.GetValues()

	for _, v := range values {
		if c.compareFunc(v, ctxValue) {
			return true
		}
	}
	return false
}
