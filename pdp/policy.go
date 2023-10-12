package pdp

import "github.com/aesoper101/pbac/pdp/condition"

// Policy - 策略
type Policy interface {
	// GetID - 返回策略的ID
	GetID() string

	// GetDescription - 返回策略的描述
	GetDescription() string

	// GetEffect - 返回策略的Effect
	GetEffect() Effect

	// GetActions - 返回策略的Actions
	GetActions() ActionSet

	// GetNotActions - 返回策略的NotActions
	GetNotActions() ActionSet

	// GetResources - 返回策略的Resources
	GetResources() []string

	// GetNotResources - 返回策略的NotResources
	GetNotResources() []string

	// GetPrincipals - 返回策略的Principals
	GetPrincipals() []string

	// GetNotPrincipals - 返回策略的NotPrincipals
	GetNotPrincipals() []string

	// GetConditions - 返回策略的Conditions
	GetConditions() condition.Functions
}
