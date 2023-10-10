package condition

import (
	"github.com/aesoper101/pbac/pdp/condition/consts"
	"github.com/aesoper101/pbac/pdp/types"
)

type numberCondition struct {
	baseKeyedCondition
	name string
	// compareFunc 比较函数, a为条件值, b为请求值
	compareFunc func(a, b interface{}) bool
}

func newNumberCondition(name string, key string, values []interface{}, compareFunc func(a, b interface{}) bool) KeyedCondition {
	return &numberCondition{
		baseKeyedCondition: baseKeyedCondition{
			key:    key,
			values: values,
		},
		name:        name,
		compareFunc: compareFunc,
	}
}

// newNumberEqualsCondition NumberEquals 请求值等于任意一个条件值。
func newNumberEqualsCondition(key string, values []interface{}) KeyedCondition {
	return newNumberCondition(consts.NumberEquals, key, values, func(a, b interface{}) bool {
		aNum, aOK := castNumber(a)
		bNum, bOK := castNumber(b)
		return aOK && bOK && aNum.Equal(bNum)
	})
}

// newNumberNotEqualsCondition NumberNotEquals 请求值不等于所有条件值。
func newNumberNotEqualsCondition(key string, values []interface{}) KeyedCondition {
	return newNumberCondition(consts.NumberNotEquals, key, values, func(a, b interface{}) bool {
		aNum, aOK := castNumber(a)
		bNum, bOK := castNumber(b)
		return aOK && bOK && !aNum.Equal(bNum)
	})
}

// newNumberLessThanCondition NumberLessThan 请求值小于所有条件值。
func newNumberLessThanCondition(key string, values []interface{}) KeyedCondition {
	return newNumberCondition(consts.NumberLessThan, key, values, func(a, b interface{}) bool {
		aNum, aOK := castNumber(a)
		bNum, bOK := castNumber(b)
		return aOK && bOK && bNum.LessThan(aNum)
	})
}

// newNumberLessThanEqualsCondition NumberLessThanEquals 请求值小于等于所有条件值。
func newNumberLessThanEqualsCondition(key string, values []interface{}) KeyedCondition {
	return newNumberCondition(consts.NumberLessThanEquals, key, values, func(a, b interface{}) bool {
		aNum, aOK := castNumber(a)
		bNum, bOK := castNumber(b)
		return aOK && bOK && bNum.LessThanOrEqual(aNum)
	})
}

// newNumberGreaterThanCondition NumberGreaterThan 请求值大于所有条件值。
func newNumberGreaterThanCondition(key string, values []interface{}) KeyedCondition {
	return newNumberCondition(consts.NumberGreaterThan, key, values, func(a, b interface{}) bool {
		aNum, aOK := castNumber(a)
		bNum, bOK := castNumber(b)
		return aOK && bOK && bNum.GreaterThan(aNum)
	})
}

// newNumberGreaterThanEqualsCondition NumberGreaterThanEquals 请求值大于等于所有条件值。
func newNumberGreaterThanEqualsCondition(key string, values []interface{}) KeyedCondition {
	return newNumberCondition(consts.NumberGreaterThanEquals, key, values, func(a, b interface{}) bool {
		aNum, aOK := castNumber(a)
		bNum, bOK := castNumber(b)
		return aOK && bOK && bNum.GreaterThanOrEqual(aNum)
	})
}

func (c *numberCondition) GetName() string {
	return c.name
}

func (c *numberCondition) Evaluate(ctxValue interface{}, _ types.EvalContextor) bool {
	values := c.GetValues()

	for _, v := range values {
		if c.compareFunc(v, ctxValue) {
			return true
		}
	}
	return false
}
