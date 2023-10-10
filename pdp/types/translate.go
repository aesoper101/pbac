package types

type Translate interface {
	// Translate 翻译
	Translate(ctx EvalContextor) (map[string]interface{}, error)
}
