package condition

type baseKeyedCondition struct {
	key    string
	values []interface{}
}

func (c *baseKeyedCondition) GetKey() string {
	return c.key
}

func (c *baseKeyedCondition) GetValues() []interface{} {
	return c.values
}
