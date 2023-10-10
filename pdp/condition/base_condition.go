package condition

import "github.com/aesoper101/pbac/pdp/types"

type baseCondition struct {
	key    string
	values []interface{}
}

func (c *baseCondition) GetKey() string {
	return c.key
}

func (c *baseCondition) GetValues() []interface{} {
	return c.values
}

func (c *baseCondition) forOr(ctxValue interface{}, _ types.EvalContextor, compareFunc func(a, b interface{}) bool) bool {
	values := c.GetValues()

	for _, v := range values {
		if compareFunc(v, ctxValue) {
			return true
		}
	}

	return false
}
