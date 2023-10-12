package condition

import (
	"fmt"
	"reflect"
	"strconv"
)

type booleanFunc struct {
	op    Operator
	k     Key
	value string
}

func (b booleanFunc) operator() Operator {
	return b.op
}

func (b booleanFunc) key() Key {
	return b.k
}

func (b booleanFunc) evaluate(values map[string][]string) bool {
	rvalues, ok := getValuesByKey(values, b.k)
	if b.op.GetPostQualifier().IsIfExistsQualifier() && !ok {
		return true
	}

	if len(rvalues) == 0 {
		return false
	}

	return b.value == rvalues[0]
}

func (b booleanFunc) String() string {
	return fmt.Sprintf("%v:%v:%v", b.op, b.k, b.value)
}

func (b booleanFunc) toMap() map[Key]ValueSet {
	return map[Key]ValueSet{
		b.k: NewValueSet(NewStringValue(b.value)),
	}
}

func (b booleanFunc) clone() Function {
	return &booleanFunc{
		op:    b.op,
		k:     b.k,
		value: b.value,
	}
}

var _ Function = (*booleanFunc)(nil)

func newBoolFunc(key Key, values ValueSet, op Operator) (Function, error) {
	if !op.IsValid() {
		return nil, fmt.Errorf("unknown operator %v", op)
	}

	if !op.GetCollectQualifier().IsZero() {
		return nil, fmt.Errorf("collect qualifier is not allowed for %v", op)
	}

	if op.GetName() != boolean {
		return nil, fmt.Errorf("unknown operator %v for %v condition", op, boolean)
	}

	if len(values) != 1 {
		return nil, fmt.Errorf("only one value is allowed for %v condition", boolean)
	}

	var value Value
	for v := range values {
		value = v
		switch v.GetType() {
		case reflect.Bool:
			if _, err := v.GetBool(); err != nil {
				return nil, err
			}
		case reflect.String:
			s, err := v.GetString()
			if err != nil {
				return nil, err
			}
			if _, err = strconv.ParseBool(s); err != nil {
				return nil, fmt.Errorf("value must be a boolean string for boolean condition")
			}
		default:
			return nil, fmt.Errorf("value must be a boolean for boolean condition")
		}
	}

	return &booleanFunc{op, key, value.String()}, nil
}

func NewBoolFunc(key Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}

	return newBoolFunc(key, values, operator)
}
