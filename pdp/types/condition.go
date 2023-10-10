package types

// Condition 条件接口
type Condition interface {
	// GetName 获取条件名称
	GetName() string

	// Evaluate 评估条件, 返回true表示条件成立
	// ctxValue: 条件上下文值, 也可以从请求上下文中获取
	// evalCtx: 评估上下文
	Evaluate(ctxValue string, evalCtx EvalContextor) bool

	// GetKey 返回条件的键
	GetKey() string

	// GetValues 返回条件的值
	GetValues() []string
}

type Conditions []Condition

func (cs Conditions) Append(c ...Condition) Conditions {
	return append(cs, c...)
}

func (cs Conditions) Evaluate(ctxValue string, evalCtx EvalContextor) bool {
	for _, c := range cs {
		if !c.Evaluate(ctxValue, evalCtx) {
			return false
		}
	}
	return true
}
