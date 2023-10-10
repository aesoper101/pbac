package ploicy

import (
	"fmt"
	"github.com/aesoper101/pbac/pdp/types"
	"github.com/samber/lo"
)

type PolicyBuilder struct {
	policy *DefaultPolicy
	errors []error
}

// NewPolicyBuilder 创建策略构建器
func NewPolicyBuilder() *PolicyBuilder {
	return &PolicyBuilder{
		policy: &DefaultPolicy{},
	}
}

func (pb *PolicyBuilder) SetID(id string) *PolicyBuilder {
	pb.policy.ID = id
	return pb
}

func (pb *PolicyBuilder) SetDescription(description string) *PolicyBuilder {
	pb.policy.Description = description
	return pb
}

func (pb *PolicyBuilder) SetEffect(effect string) *PolicyBuilder {
	if effect != types.EffectAllow && effect != types.EffectDeny {
		pb.errors = append(pb.errors, fmt.Errorf("invalid effect: %s", effect))
		return pb
	}

	pb.policy.Effect = effect
	return pb
}

func (pb *PolicyBuilder) AddResources(resources ...string) *PolicyBuilder {
	pb.policy.Resources = append(pb.policy.Resources, resources...)
	pb.policy.Resources = lo.Union(pb.policy.Resources)

	return pb
}

func (pb *PolicyBuilder) AddActions(actions ...string) *PolicyBuilder {
	pb.policy.Actions = append(pb.policy.Actions, actions...)
	pb.policy.Actions = lo.Union(pb.policy.Actions)

	return pb
}

func (pb *PolicyBuilder) AddNotResources(notResources ...string) *PolicyBuilder {
	pb.policy.NotResources = append(pb.policy.NotResources, notResources...)
	pb.policy.NotResources = lo.Union(pb.policy.NotResources)

	return pb
}

func (pb *PolicyBuilder) AddCondition(operator string, conditionKey string, conditionValue ...string) *PolicyBuilder {
	if len(pb.policy.Conditions) == 0 {
		pb.policy.Conditions = types.NewConditionMap()
	}

	kv := types.NewConditionKeyValues()
	kv.Set(conditionKey, conditionValue...)

	pb.policy.Conditions.Set(operator, kv)

	return pb
}

func (pb *PolicyBuilder) Build() (types.Policy, error) {
	if len(pb.errors) > 0 {
		return nil, pb.errors[0]
	}

	return pb.policy, nil
}
