package condition

import (
	"github.com/aesoper101/pbac/pdp/condition/consts"
	"github.com/aesoper101/pbac/pdp/types"
)

type nullCondition struct {
	baseCondition
}

func newNullCondition(key string, values []string) (types.Condition, error) {
	return &nullCondition{
		baseCondition: baseCondition{
			key:    key,
			values: values,
		},
	}, nil
}

func (c *nullCondition) GetName() string {
	return consts.Null
}

func (c *nullCondition) Evaluate(_ string, requestCtx types.EvalContextor) bool {
	values := c.GetValues()
	if len(values) == 0 {
		return false
	}

	conditionValue := values[0]
	conditionBoolValue, ok := getBoolString(conditionValue)
	if !ok {
		return false
	}

	_, exists := requestCtx.GetAttribute(c.GetKey())
	if conditionBoolValue == "true" {
		return !exists
	}

	return exists
}
