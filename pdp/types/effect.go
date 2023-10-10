package types

type Effect string

const (
	EffectAllow Effect = "Allow"
	EffectDeny  Effect = "Deny"
)

func (e Effect) String() string {
	return string(e)
}

func (e Effect) IsAllow() bool {
	return e == EffectAllow
}

func (e Effect) IsDeny() bool {
	return e == EffectDeny
}

func (e Effect) IsEqual(effect Effect) bool {
	return e == effect
}
