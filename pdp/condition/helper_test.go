package condition

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIpCondition_Evaluate(t *testing.T) {
	m := map[string]string{
		"2": "2",
	}

	v, ok := m["1"]
	k := reflect.ValueOf(&v).IsZero()
	fmt.Println(v, ok, k)
}
