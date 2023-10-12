package condition

import (
	"encoding/json"
	"strings"
)

// Key - 条件键名
type Key struct {
	name     KeyName
	variable string
}

// Is - 检查此键是否具有相同的键名
func (key Key) Is(name KeyName) bool {
	return key.name == name
}

// String - 返回键名的字符串表示形式
func (key Key) String() string {
	if key.variable != "" {
		return string(key.name) + "/" + key.variable
	}
	return string(key.name)
}

// VarName - 返回变量键名，例如"${aws:username}"
func (key Key) VarName() string {
	return key.name.VarName()
}

// Name - 返回键名，该键名是前缀"aws:"和"s3:"的值
func (key Key) Name() string {
	name := key.name.Name()
	if key.variable != "" {
		return name + "/" + key.variable
	}
	return name
}

// UnmarshalJSON - 解码JSON数据到Key
func (key *Key) UnmarshalJSON(data []byte) error {
	var keyName string
	if err := json.Unmarshal(data, &keyName); err != nil {
		return err
	}

	parsedKey, err := parseKey(keyName)
	if err != nil {
		return err
	}

	*key = parsedKey
	return nil
}

func parseKey(keyName string) (Key, error) {
	name, variable := keyName, ""
	if strings.Contains(keyName, "/") {
		tokens := strings.SplitN(keyName, "/", 2)
		name, variable = tokens[0], tokens[1]
	}

	key := Key{
		name:     KeyName(name),
		variable: variable,
	}

	return key, nil
}

// NewKey - 创建新的键
func NewKey(name KeyName, variable string) Key {
	return Key{
		name:     name,
		variable: variable,
	}
}
