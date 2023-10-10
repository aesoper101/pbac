package types

// Request 请求接口
type Request interface {
	// GetSubject 获取请求主体
	GetSubject() string

	// GetResource 获取请求资源
	GetResource() string

	// GetAction 获取请求动作
	GetAction() string

	// GetContext 获取请求上下文
	GetContext() EvalContextor
}
