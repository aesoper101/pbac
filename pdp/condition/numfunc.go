package condition

import (
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
)

type exp int

const (
	equals exp = iota + 1
	notEquals
	greaterThan
	greaterThanEquals
	lessThan
	lessThanEquals
)

type numericFunc struct {
	op    Operator
	k     Key
	value string
	e     exp
}

func (f numericFunc) operator() Operator {
	return f.op
}

func (f numericFunc) key() Key {
	return f.k
}

func (f numericFunc) evaluate(values map[string][]string) bool {
	rvalues, ok := getValuesByKey(values, f.k)
	if f.op.GetPostQualifier().IsIfExistsQualifier() && !ok {
		return true
	}

	if len(rvalues) == 0 {
		return false
	}

	rv, err := decimal.NewFromString(rvalues[0])
	if err != nil {
		return false
	}

	fv, err := decimal.NewFromString(f.value)
	if err != nil {
		return false
	}

	switch f.e {
	case equals:
		return rv.Equal(fv)
	case notEquals:
		return !rv.Equal(fv)
	case greaterThan:
		return rv.GreaterThan(fv)
	case greaterThanEquals:
		return rv.GreaterThanOrEqual(fv)
	case lessThan:
		return rv.LessThan(fv)
	case lessThanEquals:
		return rv.LessThanOrEqual(fv)
	}

	return false
}

func (f numericFunc) String() string {
	return fmt.Sprintf("%v:%v:%v", f.operator().String(), f.k, f.value)
}

func (f numericFunc) toMap() map[Key]ValueSet {
	values := NewValueSet()
	values.Add(NewStringValue(f.value))
	return map[Key]ValueSet{
		f.k: values,
	}
}

func (f numericFunc) clone() Function {
	return &numericFunc{
		op:    f.op,
		k:     f.k,
		value: f.value,
		e:     f.e,
	}
}

var _ Function = (*numericFunc)(nil)

func newNumericFunc(key Key, values ValueSet, op Operator, e exp, opName OperatorName) (Function, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("%s function requires exactly one value", op.String())
	}

	if op.IsValid() && !op.GetCollectQualifier().IsZero() {
		return nil, fmt.Errorf("unsupported collect qualifier %s for %s function", op.GetCollectQualifier().String(), op.String())
	}

	if op.GetName() != opName {
		return nil, fmt.Errorf("invalid operator %v for %v condition", op, key)
	}

	value := values.ToSlice()[0]
	if value.GetType() == reflect.Bool {
		return nil, fmt.Errorf("value must be a number or number string for %v condition", key)
	}

	if _, err := decimal.NewFromString(value.String()); err != nil {
		return nil, fmt.Errorf("value must be a number or number string for %v condition", key)
	}

	return &numericFunc{
		op:    op,
		k:     key,
		value: value.String(),
		e:     e,
	}, nil
}

func newNumericEqualsFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newNumericFunc(key, values, op, equals, numericEquals)
}

func NewNumericEqualsFunc(key Key, value Value, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}

	return newNumericEqualsFunc(key, NewValueSet(value), operator)
}

func newNumericNotEqualsFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newNumericFunc(key, values, op, notEquals, numericNotEquals)
}

func NewNumericNotEqualsFunc(key Key, value Value, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}

	return newNumericNotEqualsFunc(key, NewValueSet(value), operator)
}

func newNumericGreaterThanFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newNumericFunc(key, values, op, greaterThan, numericGreaterThan)
}

func NewNumericGreaterThanFunc(key Key, value Value, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}

	return newNumericGreaterThanFunc(key, NewValueSet(value), operator)
}

func newNumericGreaterThanEqualsFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newNumericFunc(key, values, op, greaterThanEquals, numericGreaterThanEquals)
}

func NewNumericGreaterThanEqualsFunc(key Key, value Value, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}

	return newNumericGreaterThanEqualsFunc(key, NewValueSet(value), operator)
}

func newNumericLessThanFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newNumericFunc(key, values, op, lessThan, numericLessThan)
}

func NewNumericLessThanFunc(key Key, value Value, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}

	return newNumericLessThanFunc(key, NewValueSet(value), operator)
}

func newNumericLessThanEqualsFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newNumericFunc(key, values, op, lessThanEquals, numericLessThanEquals)
}

func NewNumericLessThanEqualsFunc(key Key, value Value, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}

	return newNumericLessThanEqualsFunc(key, NewValueSet(value), operator)
}
