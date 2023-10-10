package types

// Policy 策略接口
type Policy interface {
	// GetID 获取策略ID
	GetID() string

	// GetDescription 获取策略描述
	GetDescription() string

	// AllowAccess 返回策略是否允许访问
	AllowAccess() bool

	// GetEffect 获取策略效果，可能是'Allow'或'Deny'
	GetEffect() string

	// GetResources 获取策略资源
	GetResources() []string

	// GetNotResources 获取策略非资源
	GetNotResources() []string

	// GetActions 获取策略动作
	GetActions() []string

	// GetNotActions 获取策略非动作
	GetNotActions() []string

	// GetPrincipal 获取策略主体
	GetPrincipal() PrincipalMap

	// GetNotPrincipal 获取策略非主体
	GetNotPrincipal() PrincipalMap

	// GetConditions 获取策略条件
	GetConditions() ConditionMap
}

type ConditionMap map[string]ConditionKeyValues

type ConditionKeyValues map[string][]string

func NewConditionMap() ConditionMap {
	return make(ConditionMap)
}

func (c ConditionMap) Set(key string, values ...ConditionKeyValues) {
	if _, ok := c[key]; !ok {
		c[key] = NewConditionKeyValues()
	}

	for _, v := range values {
		for k, vv := range v {
			c[key][k] = vv
		}
	}
}

func NewConditionKeyValues() ConditionKeyValues {
	return make(ConditionKeyValues)
}

func (c ConditionKeyValues) Set(key string, values ...string) {
	c[key] = values
}

func (c ConditionKeyValues) Get(key string) []string {
	return c[key]
}

type PrincipalMap map[string]PrincipalKeyValues

type PrincipalKeyValues map[string][]string
