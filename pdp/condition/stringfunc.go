package condition

import (
	"encoding/base64"
	"fmt"
	"github.com/aesoper101/pbac/internal/set"
	"github.com/aesoper101/pbac/internal/wildcard"
	"sort"
	"strings"
)

// substitute - 替换条件中的变量
func substitute(key Key, values map[string][]string) func(string) string {
	return func(v string) string {
		for k, rvalues := range values {
			if k == key.Name() && rvalues[0] != "" {
				v = strings.Replace(v, key.VarName(), rvalues[0], -1)
			}
		}
		return v
	}
}

type stringFunc struct {
	op         Operator
	k          Key
	values     set.StringSet
	ignoreCase bool
	base64     bool
	negate     bool
}

func (f stringFunc) operator() Operator {
	return f.op
}

func (f stringFunc) key() Key {
	return f.k
}

func (f stringFunc) eval(values map[string][]string) bool {
	vals, ok := getValuesByKey(values, f.k)

	// 如果该键不存在，则直接返回true
	if f.op.GetPostQualifier().IsIfExistsQualifier() && !ok {
		return true
	}

	rvalues := set.NewStringSet(vals...)
	fvalues := f.values.Clone().ApplyFunc(substitute(f.k, values))

	// 如果忽略大小写，则将所有值转换为小写
	if f.ignoreCase {
		rvalues = rvalues.ApplyFunc(strings.ToLower)
		fvalues = fvalues.ApplyFunc(strings.ToLower)
	}

	// 获取交集
	ivalues := rvalues.Intersection(fvalues)
	if f.op.GetCollectQualifier().IsForAllValuesQualifier() {
		return rvalues.IsEmpty() || rvalues.Equals(ivalues)
	}

	return !ivalues.IsEmpty()
}

func (f stringFunc) evaluate(values map[string][]string) bool {
	result := f.eval(values)
	if f.negate {
		return !result
	}
	return result
}

func (f stringFunc) String() string {
	valueStrings := f.values.ToSlice()
	sort.Strings(valueStrings)
	return fmt.Sprintf("%v:%v:%v", f.operator().String(), f.k, valueStrings)
}

func (f stringFunc) toMap() map[Key]ValueSet {
	values := NewValueSet()
	for _, value := range f.values.ToSlice() {
		if f.base64 {
			values.Add(NewStringValue(base64.StdEncoding.EncodeToString([]byte(value))))
		} else {
			values.Add(NewStringValue(value))
		}
	}

	return map[Key]ValueSet{
		f.k: values,
	}
}

func (f stringFunc) copy() stringFunc {
	return stringFunc{
		op:         f.op,
		k:          f.k,
		values:     f.values.Clone(),
		ignoreCase: f.ignoreCase,
		base64:     f.base64,
		negate:     f.negate,
	}
}

func (f stringFunc) clone() Function {
	c := f.copy()
	return &c
}

var _ Function = (*stringFunc)(nil)

// stringLikeFunc - 字符串模糊匹配函数
// 该函数用于模糊匹配字符串，支持通配符'*'和'?'，'*'匹配任意长度的字符串，'?'匹配单个字符
// 例如：'abc*'匹配'abc'、'abcd'、'abcde'等，'abc?'匹配'abca'、'abcb'、'abcc'等
type stringLikeFunc struct {
	stringFunc
}

func (f stringLikeFunc) eval(values map[string][]string) bool {
	rvalues, ok := getValuesByKey(values, f.k)
	//fmt.Println("rvalues:", rvalues, f.k.Name())
	// 如果该键不存在，则直接返回true
	if f.op.GetPostQualifier().IsIfExistsQualifier() && !ok {
		return true
	}

	fvalues := f.values.Clone().ApplyFunc(substitute(f.k, values))

	for _, v := range rvalues {
		matched := !fvalues.FuncMatch(wildcard.Match, v).IsEmpty()
		if f.op.GetCollectQualifier().IsForAllValuesQualifier() {
			if !matched {
				return false
			}
		} else if matched {
			return true
		}
	}

	return f.op.GetCollectQualifier().IsForAllValuesQualifier()
}

func (f stringLikeFunc) evaluate(values map[string][]string) bool {
	result := f.eval(values)
	if f.negate {
		return !result
	}
	return result
}

func (f stringLikeFunc) clone() Function {
	return &stringLikeFunc{stringFunc: f.copy()}
}

func newStringFunc(k Key, values ValueSet, op Operator, ignoreCase, base64, negate bool) (*stringFunc, error) {
	sset := set.NewStringSet()

	for value := range values {
		s, err := value.GetString()
		if err != nil {
			return nil, fmt.Errorf("value must be a string for %v condition", k)
		}

		sset.Add(s)
	}

	if sset.IsEmpty() {
		return nil, fmt.Errorf("value set must not be empty for %v condition", k)
	}

	if !op.IsValid() {
		return nil, fmt.Errorf("invalid operator %v for %v condition", op, k)
	}

	return &stringFunc{
		op:         op,
		k:          k,
		values:     sset,
		ignoreCase: ignoreCase,
		base64:     base64,
		negate:     negate,
	}, nil
}

func newStringFuncWithOperatorName(k Key, values ValueSet, op Operator, ignoreCase, base64, negate bool, opName OperatorName) (*stringFunc, error) {
	if op.GetName() != opName {
		return nil, fmt.Errorf("invalid operator %v for %v condition", op, k)
	}
	return newStringFunc(k, values, op, ignoreCase, base64, negate)
}

// newStringEqualsFunc - 创建字符串相等函数
func newStringEqualsFunc(k Key, values ValueSet, op Operator) (Function, error) {
	return newStringFuncWithOperatorName(k, values, op, false, false, false, stringEquals)
}

func NewStringEqualsFunc(k Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}
	return newStringEqualsFunc(k, values, operator)
}

// newStringNotEqualsFunc - 创建字符串不相等函数
func newStringNotEqualsFunc(k Key, values ValueSet, op Operator) (Function, error) {
	return newStringFuncWithOperatorName(k, values, op, false, false, true, stringNotEquals)
}

func NewStringNotEqualsFunc(k Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}
	return newStringNotEqualsFunc(k, values, operator)
}

// newStringEqualsIgnoreCaseFunc - 创建忽略大小写的字符串相等函数
func newStringEqualsIgnoreCaseFunc(k Key, values ValueSet, op Operator) (Function, error) {
	return newStringFuncWithOperatorName(k, values, op, true, false, false, stringEqualsIgnoreCase)
}

func NewStringEqualsIgnoreCaseFunc(k Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}
	return newStringEqualsIgnoreCaseFunc(k, values, operator)
}

// newStringNotEqualsIgnoreCaseFunc - 创建忽略大小写的字符串不相等函数
func newStringNotEqualsIgnoreCaseFunc(k Key, values ValueSet, op Operator) (Function, error) {
	return newStringFuncWithOperatorName(k, values, op, true, false, true, stringNotEqualsIgnoreCase)
}

func NewStringNotEqualsIgnoreCaseFunc(k Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}
	return newStringNotEqualsIgnoreCaseFunc(k, values, operator)
}

// newBinaryEqualsFunc - 创建二进制相等函数
func newBinaryEqualsFunc(k Key, values ValueSet, op Operator) (Function, error) {
	vset := NewValueSet()
	for value := range values {
		b, err := value.GetString()
		if err != nil {
			return nil, fmt.Errorf("value must be a string for %v condition", k)
		}

		data, err := base64.StdEncoding.DecodeString(b)
		if err != nil {
			return nil, err
		}
		vset.Add(NewStringValue(string(data)))
	}

	return newStringFuncWithOperatorName(k, vset, op, false, true, false, binaryEquals)
}

func NewBinaryEqualsFunc(k Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}
	return newBinaryEqualsFunc(k, values, operator)
}

// newStringLikeFunc - 创建字符串模糊匹配函数
func newStringLikeFunc(k Key, values ValueSet, op Operator) (Function, error) {
	sf, err := newStringFuncWithOperatorName(k, values, op, false, false, false, stringLike)
	if err != nil {
		return nil, err
	}

	return &stringLikeFunc{stringFunc: *sf}, nil
}

func NewStringLikeFunc(k Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}
	return newStringLikeFunc(k, values, operator)
}

// newStringNotLikeFunc - 创建字符串不模糊匹配函数
func newStringNotLikeFunc(k Key, values ValueSet, op Operator) (Function, error) {
	sf, err := newStringFuncWithOperatorName(k, values, op, false, false, true, stringNotLike)
	if err != nil {
		return nil, err
	}

	return &stringLikeFunc{stringFunc: *sf}, nil
}

func NewStringNotLikeFunc(k Key, values ValueSet, op string) (Function, error) {
	operator, err := parseOperator(op)
	if err != nil {
		return nil, err
	}
	return newStringNotLikeFunc(k, values, operator)
}
