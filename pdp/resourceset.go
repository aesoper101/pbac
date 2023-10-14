package pdp

import (
	"encoding/json"
	"fmt"
	"github.com/aesoper101/pbac/internal/set"
	"sort"
)

// ResourceSet - 资源集合
type ResourceSet map[Resource]struct{}

// NewResourceSet - 创建一个新的资源集合
func NewResourceSet(resources ...Resource) ResourceSet {
	resourceSet := ResourceSet{}
	for _, resource := range resources {
		resourceSet.Add(resource)
	}
	return resourceSet
}

// ToSlice - 返回资源集合的切片
func (resourceSet ResourceSet) ToSlice() []Resource {
	var resources []Resource
	for resource := range resourceSet {
		resources = append(resources, resource)
	}
	return resources
}

// IsEmpty - 返回资源集合是否为空
func (resourceSet ResourceSet) IsEmpty() bool {
	return len(resourceSet) == 0
}

// Add - 将资源添加到资源集合
func (resourceSet ResourceSet) Add(resource Resource) {
	resourceSet[resource] = struct{}{}
}

// Remove - 从资源集合中删除资源。如果资源不存在于集合中，则不执行任何操作。
func (resourceSet ResourceSet) Remove(resource Resource) {
	delete(resourceSet, resource)
}

// Contains - 检查资源是否在资源集合中
func (resourceSet ResourceSet) Contains(resource Resource) bool {
	_, ok := resourceSet[resource]
	return ok
}

// FuncMatch - 返回包含每个通过匹配函数的值的新集合。
// 'matchFn'应该接受集合中的元素作为第一个参数和'matchString'作为第二个参数。
// 函数可以执行任何逻辑来比较两个参数，并应返回true以接受集合中的元素以包含在输出集合中，否则将忽略该元素。
func (resourceSet ResourceSet) FuncMatch(matchFn func(Resource, Resource) bool, matchString Resource) ResourceSet {
	nset := NewResourceSet()
	for k := range resourceSet {
		if matchFn(k, matchString) {
			nset.Add(k)
		}
	}
	return nset
}

// ApplyFunc - 返回包含每个由'applyFn'处理的值的新集合。
// 'applyFn'应该接受集合中的元素作为参数，并返回一个处理后的字符串。
// 函数可以执行任何逻辑来返回一个处理后的字符串。
func (resourceSet ResourceSet) ApplyFunc(applyFn func(Resource) Resource) ResourceSet {
	nset := NewResourceSet()
	for k := range resourceSet {
		nset.Add(applyFn(k))
	}
	return nset
}

// Equals - 检查给定的资源集合是否等于当前资源集合。
func (resourceSet ResourceSet) Equals(s ResourceSet) bool {
	// 如果集合的长度不等于给定集合的长度，则集合不等于给定集合。
	if len(resourceSet) != len(s) {
		return false
	}

	// 由于两个集合的长度相等，因此检查每个元素是否相等。
	for k := range resourceSet {
		if _, ok := s[k]; !ok {
			return false
		}
	}

	return true
}

// Intersection - 返回包含两个资源集合交集的新集合
func (resourceSet ResourceSet) Intersection(sset ResourceSet) ResourceSet {
	nset := NewResourceSet()
	for k := range resourceSet {
		if _, ok := sset[k]; ok {
			nset.Add(k)
		}
	}

	return nset
}

// Difference - 返回包含两个资源集合差异的新集合
func (resourceSet ResourceSet) Difference(sset ResourceSet) ResourceSet {
	nset := NewResourceSet()
	for k := range resourceSet {
		if _, ok := sset[k]; !ok {
			nset.Add(k)
		}
	}

	return nset
}

// Clone - 克隆资源集合
func (resourceSet ResourceSet) Clone() ResourceSet {
	nset := NewResourceSet()
	for k := range resourceSet {
		nset.Add(k)
	}
	return nset
}

// MarshalJSON - 将资源集合编码为JSON数据
func (resourceSet ResourceSet) MarshalJSON() ([]byte, error) {
	if len(resourceSet) == 0 {
		return nil, fmt.Errorf("empty resources not allowed")
	}

	return json.Marshal(resourceSet.ToSlice())
}

// UnmarshalJSON - 从JSON解组资源集合
func (resourceSet *ResourceSet) UnmarshalJSON(data []byte) error {
	var sset set.StringSet
	if err := json.Unmarshal(data, &sset); err != nil {
		return err
	}

	if len(sset) == 0 {
		return fmt.Errorf("empty resources not allowed")
	}

	*resourceSet = make(ResourceSet)
	for _, s := range sset.ToSlice() {
		resourceSet.Add(Resource(s))
	}

	return nil
}

func (resourceSet ResourceSet) String() string {
	var resources []string
	for resource := range resourceSet {
		resources = append(resources, resource.String())
	}
	sort.Strings(resources)

	return fmt.Sprintf("%v", resources)
}
