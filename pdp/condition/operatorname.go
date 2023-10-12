package condition

type OperatorName string

func (on OperatorName) String() string {
	return string(on)
}

func (on OperatorName) IsValid() bool {
	_, ok := operatorNames[on]
	return ok
}

func (on OperatorName) Is(op string) bool {
	return on.String() == op
}

const (
	// names
	stringEquals              OperatorName = "StringEquals"
	stringNotEquals           OperatorName = "StringNotEquals"
	stringEqualsIgnoreCase    OperatorName = "StringEqualsIgnoreCase"
	stringNotEqualsIgnoreCase OperatorName = "StringNotEqualsIgnoreCase"
	stringLike                OperatorName = "StringLike"
	stringNotLike             OperatorName = "StringNotLike"
	binaryEquals              OperatorName = "BinaryEquals"
	ipAddress                 OperatorName = "IpAddress"
	notIPAddress              OperatorName = "NotIpAddress"
	null                      OperatorName = "Null"
	boolean                   OperatorName = "Bool"
	numericEquals             OperatorName = "NumericEquals"
	numericNotEquals          OperatorName = "NumericNotEquals"
	numericLessThan           OperatorName = "NumericLessThan"
	numericLessThanEquals     OperatorName = "NumericLessThanEquals"
	numericGreaterThan        OperatorName = "NumericGreaterThan"
	numericGreaterThanEquals  OperatorName = "NumericGreaterThanEquals"
	dateEquals                OperatorName = "DateEquals"
	dateNotEquals             OperatorName = "DateNotEquals"
	dateLessThan              OperatorName = "DateLessThan"
	dateLessThanEquals        OperatorName = "DateLessThanEquals"
	dateGreaterThan           OperatorName = "DateGreaterThan"
	dateGreaterThanEquals     OperatorName = "DateGreaterThanEquals"
)

var operatorNames = map[OperatorName]struct{}{
	stringEquals:              {},
	stringNotEquals:           {},
	stringEqualsIgnoreCase:    {},
	stringNotEqualsIgnoreCase: {},
	binaryEquals:              {},
	stringLike:                {},
	stringNotLike:             {},
	ipAddress:                 {},
	notIPAddress:              {},
	null:                      {},
	boolean:                   {},
	numericEquals:             {},
	numericNotEquals:          {},
	numericLessThan:           {},
	numericLessThanEquals:     {},
	numericGreaterThan:        {},
	numericGreaterThanEquals:  {},
	dateEquals:                {},
	dateNotEquals:             {},
	dateLessThan:              {},
	dateLessThanEquals:        {},
	dateGreaterThan:           {},
	dateGreaterThanEquals:     {},
}
