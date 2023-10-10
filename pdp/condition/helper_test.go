package condition

import (
	"fmt"
	"testing"
)

func TestIpCondition_Evaluate(t *testing.T) {

	var ms ConditionSet
	ms = make(map[string]ConditionKeyValues)
	ms["IpAddress"] = make(ConditionKeyValues)
	ms["IpAddress"]["req:IpAddress"] = []string{"3", "4"}

	for k, v := range ms {
		fmt.Printf("条件操作符=%s", k)

		for k1, v1 := range v {
			fmt.Printf("条件键=%s, 条件值=%s", k1, v1)

		}
	}

}

type ConditionSet map[string]ConditionKeyValues

type ConditionKeyValues map[string][]string
