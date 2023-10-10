package ploicy

import "github.com/aesoper101/pbac/pdp/types"

type DefaultPolicy struct {
	// ID 策略ID
	ID string `json:"Sid"`

	// Description 策略描述
	Description string `json:"Description"`

	// Effect 策略效果，可能是'Allow'或'Deny'
	Effect string `json:"Effect"`

	// Resources 策略资源
	Resources []string `json:"Resource,omitempty"`

	// Actions 策略动作
	Actions []string `json:"Action,omitempty"`

	// NotActions 策略非动作
	NotActions []string `json:"NotAction,omitempty"`

	NotResources []string `json:"NotResource,omitempty"`

	// Principal 策略主体
	Principal types.PrincipalMap `json:"Principal,omitempty"`

	// NotPrincipal 策略非主体
	NotPrincipal types.PrincipalMap `json:"NotPrincipal,omitempty"`

	// Conditions 策略条件
	Conditions types.ConditionMap `json:"Condition,omitempty"`
}

func (d *DefaultPolicy) GetID() string {
	return d.ID
}

func (d *DefaultPolicy) GetDescription() string {
	return d.Description
}

func (d *DefaultPolicy) AllowAccess() bool {
	return d.Effect == types.EffectAllow
}

func (d *DefaultPolicy) GetEffect() string {
	return d.Effect
}

func (d *DefaultPolicy) GetResources() []string {
	return d.Resources
}

func (d *DefaultPolicy) GetNotResources() []string {
	return d.NotResources
}

func (d *DefaultPolicy) GetActions() []string {
	return d.Actions
}

func (d *DefaultPolicy) GetNotActions() []string {
	return d.NotActions
}

func (d *DefaultPolicy) GetPrincipal() types.PrincipalMap {
	return d.Principal
}

func (d *DefaultPolicy) GetNotPrincipal() types.PrincipalMap {
	return d.NotPrincipal
}

func (d *DefaultPolicy) GetConditions() types.ConditionMap {
	return d.Conditions
}

var _ types.Policy = (*DefaultPolicy)(nil)
