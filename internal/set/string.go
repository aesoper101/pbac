package set

import (
	"encoding/json"
	"fmt"
	"sort"
)

type StringSet map[string]struct{}

// ToSlice - 返回字符串集合的切片
func (set StringSet) ToSlice() []string {
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

// IsEmpty - 返回字符串集合是否为空
func (set StringSet) IsEmpty() bool {
	return len(set) == 0
}

// Add - 将字符串添加到字符串集合
func (set StringSet) Add(s string) {
	set[s] = struct{}{}
}

// Remove - 从字符串集合中删除字符串。如果字符串不存在于集合中，则不执行任何操作。
func (set StringSet) Remove(s string) {
	delete(set, s)
}

// Contains - 检查字符串是否在字符串集合中
func (set StringSet) Contains(s string) bool {
	_, ok := set[s]
	return ok
}

// FuncMatch - 返回包含每个通过匹配函数的值的新集合。
// 'matchFn'应该接受集合中的元素作为第一个参数和'matchString'作为第二个参数。
// 函数可以执行任何逻辑来比较两个参数，并应返回true以接受集合中的元素以包含在输出集合中，否则将忽略该元素。
func (set StringSet) FuncMatch(matchFn func(string, string) bool, matchString string) StringSet {
	nset := NewStringSet()
	for k := range set {
		if matchFn(k, matchString) {
			nset.Add(k)
		}
	}
	return nset
}

// ApplyFunc - 返回包含每个由'applyFn'处理的值的新集合。
// 'applyFn'应该接受集合中的元素作为参数，并返回一个处理后的字符串。
// 函数可以执行任何逻辑来返回一个处理后的字符串。
func (set StringSet) ApplyFunc(applyFn func(string) string) StringSet {
	nset := NewStringSet()
	for k := range set {
		nset.Add(applyFn(k))
	}
	return nset
}

// Equals - 返回两个字符串集合是否相等
func (set StringSet) Equals(o StringSet) bool {
	if len(set) != len(o) {
		return false
	}

	for k := range set {
		if !o.Contains(k) {
			return false
		}
	}

	return true
}

// Intersection - 返回包含两个字符串集合交集的新集合
func (set StringSet) Intersection(o StringSet) StringSet {
	nset := NewStringSet()
	for k := range set {
		if o.Contains(k) {
			nset.Add(k)
		}
	}
	return nset
}

// Difference - 返回包含两个字符串集合差异的新集合
func (set StringSet) Difference(o StringSet) StringSet {
	nset := NewStringSet()
	for k := range set {
		if !o.Contains(k) {
			nset.Add(k)
		}
	}
	return nset
}

// Union - 返回包含两个字符串集合并集的新集合
func (set StringSet) Union(o StringSet) StringSet {
	nset := NewStringSet()
	for k := range set {
		nset.Add(k)
	}
	for k := range o {
		nset.Add(k)
	}
	return nset
}

// MarshalJSON - 将字符串集合转换为JSON
func (set StringSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.ToSlice())
}

// UnmarshalJSON - 从JSON数据解码字符串集合
func (set *StringSet) UnmarshalJSON(data []byte) error {
	var sl []string
	var err error
	if err = json.Unmarshal(data, &sl); err == nil {
		*set = make(StringSet)
		for _, s := range sl {
			set.Add(s)
		}
	} else {
		var s string
		if err = json.Unmarshal(data, &s); err == nil {
			*set = make(StringSet)
			set.Add(s)
		}
	}

	return err
}

func (set StringSet) String() string {
	return fmt.Sprintf("%s", set.ToSlice())
}

func (set StringSet) Clone() StringSet {
	return NewStringSet(set.ToSlice()...)
}

// NewStringSet - 创建新的字符串集合
func NewStringSet(sl ...string) StringSet {
	set := make(StringSet)
	for _, s := range sl {
		set.Add(s)
	}

	return set
}
