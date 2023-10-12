package condition

import "fmt"

// KeySet - 条件键集合
type KeySet map[Key]struct{}

// NewKeySet - 创建新的条件键集合
func NewKeySet(keys ...Key) KeySet {
	set := KeySet{}
	for _, key := range keys {
		set[key] = struct{}{}
	}
	return set
}

// Add - 将键添加到键集合
func (set KeySet) Add(key Key) {
	set[key] = struct{}{}
}

// Merge - 合并两个键集合，重复的键将被覆盖
func (set KeySet) Merge(mset KeySet) {
	for k, v := range mset {
		set[k] = v
	}
}

// ToSlice - 返回键集合的切片
func (set KeySet) ToSlice() []Key {
	keys := make([]Key, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}

// IsEmpty - 返回键集合是否为空
func (set KeySet) IsEmpty() bool {
	return len(set) == 0
}

// Match - 匹配输入键名与当前键集合
func (set KeySet) Match(key Key) bool {
	_, ok := set[key]
	if ok {
		return true
	}
	_, ok = set[key.name.ToKey()]
	return ok
}

// Difference - 返回包含两个键差异的键集合
// 示例：
//
//	keySet1 := ["one", "two", "three"]
//	keySet2 := ["two", "four", "three"]
//	keySet1.Difference(keySet2) == ["one"]
func (set KeySet) Difference(sset KeySet) KeySet {
	nset := make(KeySet)

	for k := range set {
		if !sset.Match(k) {
			nset.Add(k)
		}
	}

	return nset
}

// String - 返回键集合的字符串表示形式
func (set KeySet) String() string {
	return fmt.Sprintf("%v", set.ToSlice())
}
