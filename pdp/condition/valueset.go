package condition

import (
	"encoding/json"
	"fmt"
)

// ValueSet - 条件值集合
type ValueSet map[Value]struct{}

// NewValueSet - 创建新的条件值集合
func NewValueSet(values ...Value) ValueSet {
	set := ValueSet{}
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

// Add - 将值添加到值集合
func (set ValueSet) Add(value Value) {
	set[value] = struct{}{}
}

// Merge - 合并两个值集合，重复的值将被覆盖
func (set ValueSet) Merge(mset ValueSet) {
	for k, v := range mset {
		set[k] = v
	}
}

// ToSlice - 返回值集合的切片
func (set ValueSet) ToSlice() []Value {
	values := make([]Value, 0, len(set))
	for k := range set {
		values = append(values, k)
	}

	return values
}

// IsEmpty - 返回值集合是否为空
func (set ValueSet) IsEmpty() bool {
	return len(set) == 0
}

// Clone - 返回值集合的副本
func (set ValueSet) Clone() ValueSet {
	return NewValueSet(set.ToSlice()...)
}

// MarshalJSON - 编码值集合为JSON数据
func (set ValueSet) MarshalJSON() ([]byte, error) {
	var values []Value
	for k := range set {
		values = append(values, k)
	}

	if len(values) == 0 {
		return nil, fmt.Errorf("invalid value set %v", set)
	}

	return json.Marshal(values)
}

// UnmarshalJSON - 解码JSON数据
func (set *ValueSet) UnmarshalJSON(data []byte) error {
	var v Value
	if err := json.Unmarshal(data, &v); err == nil {
		*set = make(ValueSet)
		set.Add(v)
		return nil
	}

	var values []Value
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	if len(values) < 1 {
		return fmt.Errorf("invalid value")
	}

	*set = make(ValueSet)
	for _, v = range values {
		if _, found := (*set)[v]; found {
			return fmt.Errorf("duplicate value found '%v'", v)
		}

		set.Add(v)
	}

	return nil
}
