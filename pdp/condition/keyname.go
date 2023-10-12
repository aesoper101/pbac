package condition

import (
	"fmt"
	"strings"
)

type KeyName string

func (key KeyName) VarName() string {
	return fmt.Sprintf("${%s}", key)
}

func (key KeyName) Name() string {
	split := strings.Split(string(key), ":")
	if len(split) == 1 {
		return split[0]
	}

	return split[1]
}

func (key KeyName) Prefix() string {
	split := strings.Split(string(key), ":")
	if len(split) == 1 {
		return ""
	}

	return split[0]
}

func (key KeyName) ToKey() Key {
	return NewKey(key, "")
}
