package pdp

type Effect string

const (
	// EffectAllow - 允许
	EffectAllow Effect = "Allow"
	// EffectDeny - 拒绝
	EffectDeny Effect = "Deny"
)

// Effect - 返回Effect的字符串
func (e Effect) String() string {
	return string(e)
}

// IsAllow - 判断是否为允许
func (e Effect) IsAllow() bool {
	return e == EffectAllow
}

// IsValid - 判断是否为有效的Effect
func (e Effect) IsValid() bool {
	return e == EffectAllow || e == EffectDeny
}
