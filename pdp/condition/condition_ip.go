package condition

import (
	"github.com/aesoper101/pbac/pdp/condition/consts"
	"github.com/aesoper101/pbac/pdp/types"
	"net"
)

type ipCondition struct {
	baseCondition
	name string
	// compareFunc 比较函数, a为条件值, b为请求值
	compareFunc func(a, b string) bool
}

func newIpCondition(name string, key string, values []string, compareFunc func(a, b string) bool) (types.Condition, error) {
	return &ipCondition{
		baseCondition: baseCondition{
			key:    key,
			values: values,
		},
		name:        name,
		compareFunc: compareFunc,
	}, nil
}

// newIpAddressCondition IpAddress 请求值等于任意一个条件值。
func newIpAddressCondition(key string, values []string) (types.Condition, error) {
	return newIpCondition(consts.IpAddress, key, values, func(a, b string) bool {
		if !isString(a) || !isString(b) {
			return false
		}

		bIp := net.ParseIP(b)

		// 如果是CIDR格式的IP地址，判断是否包含在内
		_, ipNet, err := net.ParseCIDR(a)
		if err == nil {
			return ipNet.Contains(bIp)
		}

		// 如果是IP地址，判断是否相等
		return bIp.Equal(net.ParseIP(a))
	})
}

// newNotIpAddressCondition NotIpAddress 请求值不等于所有条件值。
func newNotIpAddressCondition(key string, values []string) (types.Condition, error) {
	return newIpCondition(consts.NotIpAddress, key, values, func(a, b string) bool {
		if !isString(a) || !isString(b) {
			return false
		}

		bIp := net.ParseIP(b)

		_, ipNet, err := net.ParseCIDR(a)
		if err == nil {
			return !ipNet.Contains(bIp)
		}

		return !bIp.Equal(net.ParseIP(a))
	})
}

func (c *ipCondition) GetName() string {
	return c.name
}

func (c *ipCondition) Evaluate(ctxValue string, requestCtx types.EvalContextor) bool {
	return c.forOr(ctxValue, requestCtx, c.compareFunc)
}
