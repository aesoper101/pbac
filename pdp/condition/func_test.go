package condition

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFunctions(t *testing.T) {
	funcs := NewFunctions()
	b := []byte(`{
                "StringLikeIfExists": {
                    "ec2:InstanceType": [
                        "t1.*",
                        "t2.*",
                        "m3.*"
             ],"ec1:InstanceType": [
                        "t1.*",
                        "t2.*",
                        "m3.*"
             ]}}`)
	err := funcs.UnmarshalJSON(b)
	assert.Nil(t, err)

	b, err = funcs.MarshalJSON()
	assert.Nil(t, err)

	fmt.Println(string(b))
	fmt.Println(funcs.Keys())

	ctx := make(map[string][]string)
	ctx["ec1:InstanceType"] = []string{"t1.micro"}

	evaluate := funcs.Evaluate(ctx)
	assert.True(t, evaluate)
}
