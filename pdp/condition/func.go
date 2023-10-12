package condition

import (
	"encoding/json"
	"fmt"
)

// Function - 条件接口
type Function interface {
	// Operator - 返回函数的运算符
	operator() Operator

	// Key - 返回函数的键
	key() Key

	// Evaluate - 评估函数
	evaluate(values map[string][]string) bool

	// String - 返回函数的字符串表示形式
	String() string

	// ToMap - 返回函数的map表示形式
	toMap() map[Key]ValueSet

	// Clone - 返回函数的副本
	clone() Function
}

// Functions - 函数列表
type Functions []Function

// Evaluate - 评估所有函数
func (functions Functions) Evaluate(values map[string][]string) bool {
	for _, f := range functions {
		if !f.evaluate(values) {
			return false
		}
	}
	return true
}

// Keys - 返回所有函数使用的键列表
func (functions Functions) Keys() KeySet {
	keySet := NewKeySet()
	for _, f := range functions {
		keySet.Add(f.key())
	}
	return keySet
}

// Clone - 克隆函数结构
func (functions Functions) Clone() Functions {
	var funcs []Function
	for _, f := range functions {
		funcs = append(funcs, f.clone())
	}
	return funcs
}

type FunctionBuilder func(Key, ValueSet, Operator) (Function, error)

var conditionFuncMap = map[OperatorName]FunctionBuilder{
	stringEquals:              newStringEqualsFunc,
	stringNotEquals:           newStringNotEqualsFunc,
	stringEqualsIgnoreCase:    newStringEqualsIgnoreCaseFunc,
	stringNotEqualsIgnoreCase: newStringNotEqualsIgnoreCaseFunc,
	stringLike:                newStringLikeFunc,
	stringNotLike:             newStringNotLikeFunc,
	binaryEquals:              newBinaryEqualsFunc,

	// 数字
	numericEquals:            newNumericEqualsFunc,
	numericNotEquals:         newNumericNotEqualsFunc,
	numericLessThan:          newNumericLessThanFunc,
	numericLessThanEquals:    newNumericLessThanEqualsFunc,
	numericGreaterThan:       newNumericGreaterThanFunc,
	numericGreaterThanEquals: newNumericGreaterThanEqualsFunc,

	// 布尔
	boolean: newBoolFunc,

	// Null
	null: newNullFunc,

	// 日期
	dateEquals:            newDateEqualsFunc,
	dateNotEquals:         newDateNotEqualsFunc,
	dateLessThan:          newDateLessThanFunc,
	dateLessThanEquals:    newDateLessThanEqualsFunc,
	dateGreaterThan:       newDateGreaterThanFunc,
	dateGreaterThanEquals: newDateGreaterThanEqualsFunc,

	// IP地址
	ipAddress:    newIPAddrFunc,
	notIPAddress: newNotIPAddrFunc,
}

// UnmarshalJSON - 解析JSON格式的函数
func (functions *Functions) UnmarshalJSON(data []byte) error {
	nm := make(map[string]map[string]ValueSet)
	if err := json.Unmarshal(data, &nm); err != nil {
		return err
	}

	if len(nm) == 0 {
		return fmt.Errorf("condition must not be empty")
	}

	var funcs []Function
	for nameString, args := range nm {
		operator, err := parseOperator(nameString)
		if err != nil {
			return err
		}

		for keyString, values := range args {
			key, err := parseKey(keyString)
			if err != nil {
				return err
			}

			fn, ok := conditionFuncMap[operator.GetName()]
			if !ok {
				return fmt.Errorf("function %s not found", operator.GetName())
			}

			f, err := fn(key, values, operator)
			if err != nil {
				return err
			}

			funcs = append(funcs, f)
		}
	}

	*functions = funcs
	return nil
}

// MarshalJSON - 编码函数为JSON格式
func (functions Functions) MarshalJSON() ([]byte, error) {
	nm := make(map[string]map[string]ValueSet)
	for _, f := range functions {
		op := f.operator().String()
		if nm[op] == nil {
			nm[op] = make(map[string]ValueSet)
		}

		for k, v := range f.toMap() {
			nm[op][k.String()] = v
		}
	}

	return json.Marshal(nm)
}

// NewFunctions - 创建新的函数列表
func NewFunctions(funcs ...Function) Functions {
	f := make(Functions, 0, len(funcs))
	f = append(f, funcs...)
	return f
}
