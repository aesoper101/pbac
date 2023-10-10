package ploicy

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPolicyBuilder(t *testing.T) {
	builder := NewPolicyBuilder()
	builder.SetID("TestFull").
		SetDescription("TestFull").
		SetEffect("Allow").
		AddResources("TestFull").
		AddActions("TestFull").
		AddNotResources("TestFull").
		AddCondition("StringEquals", "Key1", "Value1").
		AddCondition("StringEquals", "Key2", "Value2").
		AddCondition("StringNotEquals", "Key3", "Value3")

	policy, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, policy)

	marshal, _ := json.Marshal(policy)
	fmt.Printf("policy: %s\n", string(marshal))
}
