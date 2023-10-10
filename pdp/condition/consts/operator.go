package consts

const (
	// StringEquals 请求值与任意一个条件值相同（区分大小写）。
	StringEquals = "StringEquals"
	// StringNotEquals 请求值与任意一个条件值不相同（区分大小写）。
	StringNotEquals = "StringNotEquals"
	// StringEqualsIgnoreCase 请求值与任意一个条件值相同（不区分大小写）。
	StringEqualsIgnoreCase = "StringEqualsIgnoreCase"
	// StringNotEqualsIgnoreCase 请求值与任意一个条件值不相同（不区分大小写）。
	StringNotEqualsIgnoreCase = "StringNotEqualsIgnoreCase"
	// StringLike 请求值与任意一个条件值匹配（区分大小写,正则表达式仅支持*和?）。
	StringLike = "StringLike"
	// StringNotLike 请求值与任意一个条件值不匹配（区分大小写,正则表达式仅支持*和?）。
	StringNotLike = "StringNotLike"

	// NumberEquals 请求值与任意一个条件值相同。
	NumberEquals = "NumericEquals"
	// NumberNotEquals 请求值与任意一个条件值不相同。
	NumberNotEquals = "NumericNotEquals"
	// NumberLessThan 请求值小于任意一个条件值。
	NumberLessThan = "NumericLessThan"
	// NumberLessThanEquals 请求值小于或等于任意一个条件值。
	NumberLessThanEquals = "NumericLessThanEquals"
	// NumberGreaterThan 请求值大于任意一个条件值。
	NumberGreaterThan = "NumericGreaterThan"
	// NumberGreaterThanEquals 请求值大于或等于任意一个条件值。
	NumberGreaterThanEquals = "NumericGreaterThanEquals"

	// DateEquals 请求值与任意一个条件值相同。
	DateEquals = "DateEquals"
	// DateNotEquals 请求值与任意一个条件值不相同。
	DateNotEquals = "DateNotEquals"
	// DateLessThan 请求值小于任意一个条件值。
	DateLessThan = "DateLessThan"
	// DateLessThanEquals 请求值小于或等于任意一个条件值。
	DateLessThanEquals = "DateLessThanEquals"
	// DateGreaterThan 请求值大于任意一个条件值。
	DateGreaterThan = "DateGreaterThan"
	// DateGreaterThanEquals 请求值大于或等于任意一个条件值。
	DateGreaterThanEquals = "DateGreaterThanEquals"

	// Bool 请求值与任意一个条件值相同。
	Bool = "Bool"

	// IpAddress 请求值与任意一个条件值相同。
	IpAddress = "IpAddress"
	// NotIpAddress 请求值与任意一个条件值不相同。
	NotIpAddress = "NotIpAddress"

	// Null 要求请求值不存在或者值为null
	Null = "Null"
)

const (
	Or  = "Or"
	And = "And"
)
