package condition

import (
	"github.com/aesoper101/pbac/pdp/condition/consts"
	"github.com/aesoper101/pbac/pdp/types"
)

type boolCondition struct {
	baseKeyedCondition
}

func newBoolCondition(key string, values []interface{}) KeyedCondition {
	return &boolCondition{
		baseKeyedCondition: baseKeyedCondition{
			key:    key,
			values: values,
		},
	}
}

func (c *boolCondition) GetName() string {
	return consts.Bool
}

func (c *boolCondition) Evaluate(ctxValue interface{}, requestCtx types.EvalContextor) bool {
	values := c.GetValues()
	if len(values) == 0 {
		return false
	}

	a, aOK := getBoolString(values[0])
	b, bOK := getBoolString(ctxValue)
	if !aOK || !bOK {
		return false
	}

	return b == a
}
