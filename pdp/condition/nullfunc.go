package condition

import (
	"fmt"
	"reflect"
	"strconv"
)

type nullFunc struct {
	op    Operator
	k     Key
	value bool
}

func (n nullFunc) operator() Operator {
	return n.op
}

func (n nullFunc) key() Key {
	return n.k
}

func (n nullFunc) evaluate(values map[string][]string) bool {
	rvalues, _ := getValuesByKey(values, n.k)
	if n.value {
		return len(rvalues) == 0
	}

	return len(rvalues) != 0
}

func (n nullFunc) String() string {
	return fmt.Sprintf("%v:%v:%v", n.op, n.k, n.value)
}

func (n nullFunc) toMap() map[Key]ValueSet {
	return map[Key]ValueSet{
		n.k: NewValueSet(NewBoolValue(n.value)),
	}
}

func (n nullFunc) clone() Function {
	return &nullFunc{
		op:    n.op,
		k:     n.k,
		value: n.value,
	}
}

var _ Function = (*nullFunc)(nil)

func newNullFunc(key Key, values ValueSet, op Operator) (Function, error) {
	if !op.IsValid() {
		return nil, fmt.Errorf("unknown operator %v", op)
	}

	if !op.GetCollectQualifier().IsZero() {
		return nil, fmt.Errorf("collect qualifier is not allowed for null condition")
	}

	if !op.GetPostQualifier().IsZero() {
		return nil, fmt.Errorf("post qualifier is not allowed for null condition")
	}

	if op.GetName() != null {
		return nil, fmt.Errorf("unknown operator %v for null condition", op)
	}

	if len(values) != 1 {
		return nil, fmt.Errorf("only one value is allowed for null condition")
	}

	var value bool
	for v := range values {
		switch v.GetType() {
		case reflect.Bool:
			value, _ = v.GetBool()
		case reflect.String:
			var err error
			s, _ := v.GetString()
			if value, err = strconv.ParseBool(s); err != nil {
				return nil, fmt.Errorf("value must be a boolean string for Null condition")
			}
		default:
			return nil, fmt.Errorf("value must be a boolean for Null condition")
		}
	}

	return &nullFunc{
		op:    op,
		k:     key,
		value: value,
	}, nil
}

func NewNullFunc(key Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}

	return newNullFunc(key, values, operator)
}
