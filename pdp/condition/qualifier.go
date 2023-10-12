package condition

import (
	"encoding/json"
	"fmt"
)

// Qualifier 条件集合运算符
// type Qualifier string
const (
	// forAnyValue  任意条件满足
	forAnyValue = "ForAnyValue"
	// forAllValues 所有条件满足
	forAllValues = "ForAllValues"

	// ZERO 无条件集合运算符
	zero = ""

	// IfExists IfExists后置条件集合运算符
	IfExists = "IfExists"
)

// CollectQualifier 从条件集合中收集条件集合运算符
type CollectQualifier struct {
	qualifier string
}

func (cq CollectQualifier) IsValid() bool {
	switch cq.qualifier {
	case forAnyValue, forAllValues, zero:
		return true
	default:
		return false
	}
}

func (cq CollectQualifier) String() string {
	return cq.qualifier
}

func (cq CollectQualifier) IsForAnyValueQualifier() bool {
	return cq.qualifier == forAnyValue
}

func (cq CollectQualifier) IsForAllValuesQualifier() bool {
	return cq.qualifier == forAllValues
}

func (cq CollectQualifier) IsZero() bool {
	return cq.qualifier == zero
}

func (cq CollectQualifier) MarshalJSON() ([]byte, error) {
	if !cq.IsValid() {
		return nil, fmt.Errorf("invalid qualifier %v", cq)
	}

	return json.Marshal(cq.String())
}

func (cq *CollectQualifier) UnmarshalJSON(data []byte) error {
	var qualifier string
	if err := json.Unmarshal(data, &qualifier); err != nil {
		return err
	}

	cq.qualifier = qualifier
	return nil
}

type PostQualifier struct {
	qualifier string
}

func (pq PostQualifier) IsValid() bool {
	switch pq.qualifier {
	case IfExists, zero:
		return true
	default:
		return false
	}
}

func (pq PostQualifier) String() string {
	return pq.qualifier
}

func (pq PostQualifier) IsIfExistsQualifier() bool {
	return pq.qualifier == IfExists
}

func (pq PostQualifier) IsZero() bool {
	return pq.qualifier == zero
}

func (pq PostQualifier) MarshalJSON() ([]byte, error) {
	if !pq.IsValid() {
		return nil, fmt.Errorf("invalid qualifier %v", pq)
	}

	return json.Marshal(pq.String())
}

func (pq *PostQualifier) UnmarshalJSON(data []byte) error {
	var qualifier string
	if err := json.Unmarshal(data, &qualifier); err != nil {
		return err
	}

	pq.qualifier = qualifier
	return nil
}

func NewCollectQualifier(qualifier string) CollectQualifier {
	return CollectQualifier{qualifier: qualifier}
}

func NewPostQualifier(qualifier string) PostQualifier {
	return PostQualifier{qualifier: qualifier}
}
