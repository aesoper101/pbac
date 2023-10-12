package condition

import (
	"fmt"
	"net"
	"sort"
	"strings"
)

type ipaddrFunc struct {
	op     Operator
	k      Key
	values []*net.IPNet
	negate bool
}

func (f ipaddrFunc) operator() Operator {
	return f.op
}

func (f ipaddrFunc) key() Key {
	return f.k
}

func (f ipaddrFunc) eval(values map[string][]string) bool {
	rvalues, ok := getValuesByKey(values, f.k)
	if f.op.GetPostQualifier().IsIfExistsQualifier() && !ok {
		return true
	}

	if len(rvalues) == 0 {
		return false
	}

	var IPs []net.IP
	for _, s := range rvalues {
		IP := net.ParseIP(s)
		if IP == nil {
			panic(fmt.Errorf("invalid IP address '%v'", s))
		}

		IPs = append(IPs, IP)
	}

	for _, IP := range IPs {
		for _, IPNet := range f.values {
			if IPNet.Contains(IP) {
				return true
			}
		}
	}

	return false
}

func (f ipaddrFunc) evaluate(values map[string][]string) bool {
	result := f.eval(values)
	if f.negate {
		return !result
	}
	return result
}

func (f ipaddrFunc) String() string {
	var valueStrings []string
	for _, value := range f.values {
		valueStrings = append(valueStrings, value.String())
	}
	sort.Strings(valueStrings)

	return fmt.Sprintf("%v:%v:%v", f.op, f.k, valueStrings)
}

func (f ipaddrFunc) toMap() map[Key]ValueSet {
	values := make(ValueSet)
	for _, value := range f.values {
		values[NewStringValue(value.String())] = struct{}{}
	}

	return map[Key]ValueSet{
		f.k: values,
	}
}

func (f ipaddrFunc) clone() Function {
	var values []*net.IPNet
	for _, value := range f.values {
		_, IPNet, _ := net.ParseCIDR(value.String())
		values = append(values, IPNet)
	}
	return &ipaddrFunc{
		op:     f.op,
		k:      f.k,
		values: values,
		negate: f.negate,
	}
}

var _ Function = (*ipaddrFunc)(nil)

func newIPAddrFuncWithKey(key Key, values ValueSet, op Operator, negate bool, opName OperatorName) (Function, error) {
	if !op.IsValid() {
		return nil, fmt.Errorf("unknown operator %v", op)
	}

	if !op.GetCollectQualifier().IsZero() {
		return nil, fmt.Errorf("operator %v cannot have collect qualifier", op)
	}

	IPNets, err := valuesToIPNets(opName.String(), values)
	if err != nil {
		return nil, err
	}

	return &ipaddrFunc{
		op:     op,
		k:      key,
		values: IPNets,
		negate: negate,
	}, nil
}

func newIPAddrFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newIPAddrFuncWithKey(key, values, op, false, ipAddress)
}

func newNotIPAddrFunc(key Key, values ValueSet, op Operator) (Function, error) {
	return newIPAddrFuncWithKey(key, values, op, true, notIPAddress)
}

func valuesToIPNets(n string, values ValueSet) ([]*net.IPNet, error) {
	var IPNets []*net.IPNet
	for v := range values {
		s, err := v.GetString()
		if err != nil {
			return nil, fmt.Errorf("value %v must be string representation of CIDR for %v condition", v, n)
		}

		// If you specify an IP address without the associated routing prefix, IAM uses the default prefix value of /32.
		// https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html#Conditions_IPAddress
		if strings.IndexByte(s, '/') == -1 {
			s += "/32"
		}

		var IPNet *net.IPNet
		_, IPNet, err = net.ParseCIDR(s)
		if err != nil {
			return nil, fmt.Errorf("value %v must be CIDR string for %v condition", s, n)
		}

		IPNets = append(IPNets, IPNet)
	}

	return IPNets, nil
}
