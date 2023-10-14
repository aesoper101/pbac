package pdp

// Resource - 资源
type Resource string

// Resource - 返回Resource的字符串
func (r Resource) String() string {
	return string(r)
}

// Match - 检查资源是否匹配
//func (r Resource) Match(s Resource, conditionValues map[string][]string) bool {
//	// 替换中 r 中的变量
//
//}

// 检查r中是否存在变量

func NewResource(resource string) Resource {
	return Resource(resource)
}
