package types

//go:generate mockgen -destination=../mock/mock_context.go -package=mock github.com/aesoper101/pbac/pdp/types EvalContextor

// EvalContextor 是一个接口，用于获取条件上下文属性
type EvalContextor interface {
	GetAttribute(key string) (interface{}, bool)
}
