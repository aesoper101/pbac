package pdp

import "github.com/aesoper101/pbac/internal/wildcard"

type Action string

// Action - 返回Action的字符串
func (a Action) String() string {
	return string(a)
}

// Match - 检查动作是否匹配
func (a Action) Match(s Action) bool {
	return wildcard.Match(string(a), string(s))
}

func NewAction(action string) Action {
	return Action(action)
}
