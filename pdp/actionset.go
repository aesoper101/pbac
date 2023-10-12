package pdp

import (
	"encoding/json"
	"fmt"
	"github.com/aesoper101/pbac/internal/set"
)

type ActionSet map[Action]struct{}

// ToSlice - 返回动作集合的切片
func (actionSet ActionSet) ToSlice() []Action {
	keys := make([]Action, 0, len(actionSet))
	for k := range actionSet {
		keys = append(keys, k)
	}

	return keys
}

// IsEmpty - 返回动作集合是否为空
func (actionSet ActionSet) IsEmpty() bool {
	return len(actionSet) == 0
}

// Add - 将动作添加到动作集合
func (actionSet ActionSet) Add(s Action) {
	actionSet[s] = struct{}{}
}

// Remove - 从动作集合中删除动作。如果动作不存在于集合中，则不执行任何操作。
func (actionSet ActionSet) Remove(s Action) {
	delete(actionSet, s)
}

// Contains - 检查动作是否在动作集合中
func (actionSet ActionSet) Contains(s Action) bool {
	_, ok := actionSet[s]
	return ok
}

// FuncMatch - 返回包含每个通过匹配函数的值的新集合。
// 'matchFn'应该接受集合中的元素作为第一个参数和'matchString'作为第二个参数。
// 函数可以执行任何逻辑来比较两个参数，并应返回true以接受集合中的元素以包含在输出集合中，否则将忽略该元素。
func (actionSet ActionSet) FuncMatch(matchFn func(Action, Action) bool, matchString Action) ActionSet {
	nset := NewActionSet()
	for k := range actionSet {
		if matchFn(k, matchString) {
			nset.Add(k)
		}
	}
	return nset
}

// ApplyFunc - 返回包含每个由'applyFn'处理的值的新集合。
// 'applyFn'应该接受集合中的元素作为参数，并返回一个处理后的字符串。
// 函数可以执行任何逻辑来返回一个处理后的字符串。
func (actionSet ActionSet) ApplyFunc(applyFn func(Action) Action) ActionSet {
	nset := NewActionSet()
	for k := range actionSet {
		nset.Add(applyFn(k))
	}
	return nset
}

// Equals - 检查给定的动作集合是否等于当前动作集合。
func (actionSet ActionSet) Equals(s ActionSet) bool {
	// 如果集合的长度不等于给定集合的长度，则集合不等于给定集合。
	if len(actionSet) != len(s) {
		return false
	}

	// 由于两个集合的长度相等，因此检查每个元素是否相等。
	for k := range actionSet {
		if _, ok := s[k]; !ok {
			return false
		}
	}

	return true
}

// Intersection - 返回两个动作集合交集的新集合
func (actionSet ActionSet) Intersection(sset ActionSet) ActionSet {
	nset := NewActionSet()
	for k := range actionSet {
		if _, ok := sset[k]; ok {
			nset.Add(k)
		}
	}

	return nset
}

// Difference - 返回两个动作集合差异的新集合
func (actionSet ActionSet) Difference(sset ActionSet) ActionSet {
	nset := NewActionSet()
	for k := range actionSet {
		if _, ok := sset[k]; !ok {
			nset.Add(k)
		}
	}

	return nset
}

// Union - 返回两个动作集合并集的新集合
func (actionSet ActionSet) Union(sset ActionSet) ActionSet {
	nset := NewActionSet()
	for k := range actionSet {
		nset.Add(k)
	}
	for k := range sset {
		nset.Add(k)
	}

	return nset
}

// UnmarshalJSON - 从JSON解组动作集合
func (actionSet *ActionSet) UnmarshalJSON(data []byte) error {
	var sset set.StringSet
	if err := json.Unmarshal(data, &sset); err != nil {
		return err
	}

	if len(sset) == 0 {
		return fmt.Errorf("empty actions not allowed")
	}

	*actionSet = make(ActionSet)
	for _, s := range sset.ToSlice() {
		actionSet.Add(Action(s))
	}

	return nil
}

// MarshalJSON - 将动作集合转换为JSON
func (actionSet ActionSet) MarshalJSON() ([]byte, error) {
	if len(actionSet) == 0 {
		return nil, fmt.Errorf("empty actions not allowed")
	}

	return json.Marshal(actionSet.ToSlice())
}

// NewActionSet - 返回新的动作集合
func NewActionSet(actions ...Action) ActionSet {
	actionSet := make(ActionSet)
	for _, action := range actions {
		actionSet.Add(action)
	}
	return actionSet
}
