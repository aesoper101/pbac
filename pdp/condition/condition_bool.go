package condition

import (
	"github.com/aesoper101/pbac/pdp/condition/consts"
	"github.com/aesoper101/pbac/pdp/types"
)

type boolCondition struct {
	baseCondition
}

func newBoolCondition(key string, values []string) (types.Condition, error) {
	return &boolCondition{
		baseCondition: baseCondition{
			key:    key,
			values: values,
		},
	}, nil
}

func (c *boolCondition) GetName() string {
	return consts.Bool
}

func (c *boolCondition) Evaluate(ctxValue string, _ types.EvalContextor) bool {
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
