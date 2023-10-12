package condition

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"reflect"
)

type dateFunc struct {
	op    Operator
	k     Key
	value string
	e     exp
}

func (d dateFunc) operator() Operator {
	return d.op
}

func (d dateFunc) key() Key {
	return d.k
}

func (d dateFunc) evaluate(values map[string][]string) bool {
	rvalues, ok := getValuesByKey(values, d.k)
	if d.op.GetPostQualifier().IsIfExistsQualifier() && !ok {
		return true
	}

	if len(rvalues) == 0 {
		return false
	}

	rdate := carbon.Parse(rvalues[0])
	if rdate.IsValid() {
		return false
	}

	fdate := carbon.Parse(d.value)

	switch d.e {
	case equals:
		return rdate.Eq(fdate)
	case notEquals:
		return !rdate.Eq(fdate)
	case greaterThan:
		return rdate.Gt(fdate)
	case greaterThanEquals:
		return rdate.Gte(fdate)
	case lessThan:
		return rdate.Lt(fdate)
	case lessThanEquals:
		return rdate.Lte(fdate)
	}

	return false
}

func (d dateFunc) String() string {
	return fmt.Sprintf("%v:%v:%v", d.op, d.k, d.value)
}

func (d dateFunc) toMap() map[Key]ValueSet {
	return map[Key]ValueSet{
		d.k: NewValueSet(NewStringValue(d.value)),
	}
}

func (d dateFunc) clone() Function {
	return &dateFunc{
		op:    d.op,
		k:     d.k,
		value: d.value,
	}
}

var _ Function = (*dateFunc)(nil)

func newDateFunc(key Key, values ValueSet, op Operator, e exp, opName OperatorName) (Function, error) {
	if !op.IsValid() {
		return nil, fmt.Errorf("unknown operator %v", op)
	}

	if !op.GetCollectQualifier().IsZero() {
		return nil, fmt.Errorf("collect qualifier is not allowed for date condition")
	}

	if op.GetName() != opName {
		return nil, fmt.Errorf("unknown operator %v for date condition", op)
	}

	if len(values) == 0 {
		return nil, fmt.Errorf("value is required for date condition")
	}

	if len(values) != 1 {
		return nil, fmt.Errorf("only one value is allowed for date condition")
	}

	var value Value
	for v := range values {
		value = v
		switch v.GetType() {
		case reflect.String:
			s, err := v.GetString()
			if err != nil {
				return nil, err
			}

			c := carbon.Parse(s)
			if !c.IsValid() {
				return nil, fmt.Errorf("value %s must be a RFC3339 time string for date condition", s)
			}

			return &dateFunc{
				op:    op,
				k:     key,
				value: s,
				e:     e,
			}, nil
		default:
			return nil, fmt.Errorf("value must be a string for date condition")
		}
	}

	return nil, fmt.Errorf("invalid value type %v", value.GetType())
}

func newDateEqualsFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newDateFunc(key, values, op, equals, dateEquals)
}

func newDateNotEqualsFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newDateFunc(key, values, op, notEquals, dateNotEquals)
}

func newDateGreaterThanFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newDateFunc(key, values, op, greaterThan, dateGreaterThan)
}

func newDateGreaterThanEqualsFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newDateFunc(key, values, op, greaterThanEquals, dateGreaterThanEquals)
}

func newDateLessThanFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newDateFunc(key, values, op, lessThan, dateLessThan)
}

func newDateLessThanEqualsFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newDateFunc(key, values, op, lessThanEquals, dateLessThanEquals)
}

func NewDateEqualsFunc(key Key, set ValueSet, op Operator) (Function, error) {
	return newDateEqualsFunc(key, set, op)
}

func NewDateNotEqualsFunc(key Key, set ValueSet, op Operator) (Function, error) {
	return newDateNotEqualsFunc(key, set, op)
}

func NewDateGreaterThanFunc(key Key, set ValueSet, op Operator) (Function, error) {
	return newDateGreaterThanFunc(key, set, op)
}

func NewDateGreaterThanEqualsFunc(key Key, set ValueSet, op Operator) (Function, error) {
	return newDateGreaterThanEqualsFunc(key, set, op)
}

func NewDateLessThanFunc(key Key, set ValueSet, op Operator) (Function, error) {
	return newDateLessThanFunc(key, set, op)
}

func NewDateLessThanEqualsFunc(key Key, set ValueSet, op Operator) (Function, error) {
	return newDateLessThanEqualsFunc(key, set, op)
}
