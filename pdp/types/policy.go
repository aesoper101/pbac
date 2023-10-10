package types

// Policy 策略接口
type Policy interface {
	// GetID 获取策略ID
	GetID() string

	// GetDescription 获取策略描述
	GetDescription() string

	// GetSubjects 获取策略主体
	GetSubjects() []string

	// AllowAccess 返回策略是否允许访问
	AllowAccess() bool

	// GetEffect 获取策略效果，可能是'Allow'或'Deny'
	GetEffect() string

	// GetResources 获取策略资源
	GetResources() []string

	// GetActions 获取策略动作
	GetActions() []string

	// GetConditions 获取策略条件
	GetConditions()
}
