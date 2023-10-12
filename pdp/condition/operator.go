package condition

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Operator - 条件运算符
type Operator struct {
	raw              string
	name             OperatorName
	postQualifier    PostQualifier
	collectQualifier CollectQualifier
}

// GetName - 获取条件运算符名称
func (o Operator) GetName() OperatorName {
	return o.name
}

// GetPostQualifier - 获取后置条件集合运算符
func (o Operator) GetPostQualifier() PostQualifier {
	return o.postQualifier
}

// GetCollectQualifier - 获取条件集合运算符
func (o Operator) GetCollectQualifier() CollectQualifier {
	return o.collectQualifier
}

// IsValid - 判断条件运算符是否有效
func (o Operator) IsValid() bool {
	return o.postQualifier.IsValid() && o.collectQualifier.IsValid()
}

func (o Operator) String() string {
	if !o.IsValid() {
		return ""
	}

	return o.raw
}

// UnmarshalJSON - 解码JSON数据到Operator
func (o *Operator) UnmarshalJSON(data []byte) error {
	var operator string
	if err := json.Unmarshal(data, &operator); err != nil {
		return err
	}

	parsedOperator, err := parseOperator(operator)
	if err != nil {
		return err
	}

	if !parsedOperator.IsValid() {
		return fmt.Errorf("invalid operator %v", parsedOperator)
	}

	*o = parsedOperator
	return nil
}

// parseOperator 解析成Operator
func parseOperator(name string) (Operator, error) {
	cq := NewCollectQualifier(zero)
	pq := NewPostQualifier(zero)
	op := name
	if strings.Contains(name, ":") {
		splitN := strings.SplitN(name, ":", 2)
		cq = NewCollectQualifier(splitN[0])
		op = strings.Replace(splitN[1], IfExists, "", 1)
	} else {
		op = strings.Replace(name, IfExists, "", 1)
	}

	if strings.Contains(name, IfExists) {
		pq = NewPostQualifier(IfExists)
	}

	if !pq.IsValid() {
		return Operator{}, fmt.Errorf("invalid post qualifier '%v'", name)
	}

	if !cq.IsValid() {
		return Operator{}, fmt.Errorf("invalid collect qualifier '%v'", name)
	}

	opName := OperatorName(op)
	if !opName.IsValid() {
		return Operator{}, fmt.Errorf("invalid operator '%v'", name)
	}

	return Operator{
		raw:              name,
		postQualifier:    pq,
		collectQualifier: cq,
		name:             opName,
	}, nil
}
