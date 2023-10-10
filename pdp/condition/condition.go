package condition

import (
	"github.com/aesoper101/pbac/pdp/condition/consts"
	"github.com/aesoper101/pbac/pdp/types"
	"reflect"
)

type conditionFunc func(key string, values []interface{}) (types.Condition, error)

var conditionFactories map[string]conditionFunc

func init() {
	conditionFactories = map[string]conditionFunc{
		// 布尔条件
		consts.Bool: newBoolCondition,

		// 空条件
		consts.Null: newNullCondition,

		// 字符串条件
		consts.StringEquals:              newStringEqualsCondition,
		consts.StringNotEquals:           newStringNotEqualsCondition,
		consts.StringEqualsIgnoreCase:    newStringEqualsIgnoreCaseCondition,
		consts.StringNotEqualsIgnoreCase: newStringNotEqualsIgnoreCaseCondition,
		consts.StringLike:                newStringLikeCondition,
		consts.StringNotLike:             newStringNotLikeCondition,

		// 日期条件
		consts.DateEquals:            newDateEqualsCondition,
		consts.DateNotEquals:         newDateNotEqualsCondition,
		consts.DateLessThan:          newDateLessThanCondition,
		consts.DateLessThanEquals:    newDateLessThanEqualsCondition,
		consts.DateGreaterThan:       newDateGreaterThanCondition,
		consts.DateGreaterThanEquals: newDateGreaterThanEqualsCondition,

		// IP地址条件
		consts.IpAddress:    newIpAddressCondition,
		consts.NotIpAddress: newNotIpAddressCondition,

		// 数字条件
		consts.NumberEquals:            newNumberEqualsCondition,
		consts.NumberNotEquals:         newNumberNotEqualsCondition,
		consts.NumberLessThan:          newNumberLessThanCondition,
		consts.NumberLessThanEquals:    newNumberLessThanEqualsCondition,
		consts.NumberGreaterThan:       newNumberGreaterThanCondition,
		consts.NumberGreaterThanEquals: newNumberGreaterThanEqualsCondition,
	}

	copyConditionFactories := make(map[string]conditionFunc)
	for k, v := range conditionFactories {
		copyConditionFactories[k] = v
	}

	// 添加IfExists条件
	for k, op := range copyConditionFactories {
		conditionFactories[k+consts.QualifierIfExists] = newIfExistsCondition(op)
		copyConditionFactories[k+consts.QualifierIfExists] = newIfExistsCondition(op)
	}

	// 添加ForAllValues条件
	for k, op := range copyConditionFactories {
		conditionFactories[consts.QualifierForAllValues+":"+k] = newForAllValuesCondition(op)
	}

	// 添加ForAnyValue条件
	for k, op := range copyConditionFactories {
		conditionFactories[consts.QualifierForAnyValue+":"+k] = newForAnyValueCondition(op)
	}

	// 删除copyConditionFactories
	copyConditionFactories = nil
}

type IfExistsCondition struct {
	condition types.Condition
}

func newIfExistsCondition(fn conditionFunc) conditionFunc {
	return func(key string, values []interface{}) (types.Condition, error) {
		condition, err := fn(key, values)
		if err != nil {
			return nil, err
		}
		return &IfExistsCondition{
			condition: condition,
		}, nil
	}
}

func (c *IfExistsCondition) GetName() string {
	return c.condition.GetName() + consts.QualifierIfExists
}

func (c *IfExistsCondition) Evaluate(ctxValue interface{}, requestCtx types.EvalContextor) bool {
	key := c.condition.GetKey()
	_, exists := requestCtx.GetAttribute(key)
	if !exists {
		return true
	}

	return c.condition.Evaluate(ctxValue, requestCtx)
}

func (c *IfExistsCondition) GetKey() string {
	return c.condition.GetKey()
}

func (c *IfExistsCondition) GetValues() []interface{} {
	return c.condition.GetValues()
}

type ForAllValuesCondition struct {
	condition types.Condition
}

func newForAllValuesCondition(fn conditionFunc) conditionFunc {
	return func(key string, values []interface{}) (types.Condition, error) {
		condition, err := fn(key, values)
		if err != nil {
			return nil, err
		}
		return &ForAllValuesCondition{
			condition: condition,
		}, nil
	}
}

func (c *ForAllValuesCondition) GetName() string {
	return consts.QualifierForAllValues + ":" + c.condition.GetName()
}

func (c *ForAllValuesCondition) Evaluate(ctxValue interface{}, requestCtx types.EvalContextor) bool {
	valOf := reflect.ValueOf(ctxValue)
	switch valOf.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < valOf.Len(); i++ {
			if !c.condition.Evaluate(valOf.Index(i).Interface(), requestCtx) {
				return false
			}
		}

		return true
	default:
		return c.condition.Evaluate(ctxValue, requestCtx)
	}
}

func (c *ForAllValuesCondition) GetKey() string {
	return c.condition.GetKey()
}

func (c *ForAllValuesCondition) GetValues() []interface{} {
	return c.condition.GetValues()
}

type ForAnyValueCondition struct {
	condition types.Condition
}

func newForAnyValueCondition(fn conditionFunc) conditionFunc {
	return func(key string, values []interface{}) (types.Condition, error) {
		condition, err := fn(key, values)
		if err != nil {
			return nil, err
		}
		return &ForAnyValueCondition{
			condition: condition,
		}, nil
	}
}

func (c *ForAnyValueCondition) GetName() string {
	return consts.QualifierForAnyValue + ":" + c.condition.GetName()
}

func (c *ForAnyValueCondition) Evaluate(ctxValue interface{}, requestCtx types.EvalContextor) bool {
	valOf := reflect.ValueOf(ctxValue)
	switch valOf.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < valOf.Len(); i++ {
			if c.condition.Evaluate(valOf.Index(i).Interface(), requestCtx) {
				return true
			}
		}

		return false
	default:
		return c.condition.Evaluate(ctxValue, requestCtx)
	}
}

func (c *ForAnyValueCondition) GetKey() string {
	return c.condition.GetKey()
}

func (c *ForAnyValueCondition) GetValues() []interface{} {
	return c.condition.GetValues()
}
