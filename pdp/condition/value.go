package condition

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Value struct {
	t reflect.Kind
	s string
	i int
	b bool
	f float64
}

func (v Value) GetBool() (bool, error) {
	var err error

	if v.t != reflect.Bool {
		err = fmt.Errorf("not a bool Value")
	}

	return v.b, err
}

func (v Value) GetFloat() (float64, error) {
	var err error

	if v.t != reflect.Float64 {
		err = fmt.Errorf("not a float Value")
	}

	return v.f, err
}

func (v Value) GetInt() (int, error) {
	var err error

	if v.t != reflect.Int {
		err = fmt.Errorf("not a int Value")
	}

	return v.i, err
}

func (v Value) GetString() (string, error) {
	var err error

	if v.t != reflect.String {
		err = fmt.Errorf("not a string Value")
	}

	return v.s, err
}

func (v Value) GetType() reflect.Kind {
	return v.t
}

func (v Value) MarshalJSON() ([]byte, error) {
	switch v.t {
	case reflect.String:
		return json.Marshal(v.s)
	case reflect.Int:
		return json.Marshal(v.i)
	case reflect.Bool:
		return json.Marshal(v.b)
	case reflect.Float64:
		return json.Marshal(v.f)
	}
	return nil, fmt.Errorf("unknown type %v", v.t)
}

func (v *Value) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		v.t = reflect.String
		v.s = s
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		v.t = reflect.Int
		v.i = i
		return nil
	}

	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		v.t = reflect.Bool
		v.b = b
		return nil
	}

	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		v.t = reflect.Float64
		v.f = f
		return nil
	}

	return fmt.Errorf("invalid value %s", string(data))
}

func (v Value) String() string {
	switch v.t {
	case reflect.String:
		return v.s
	case reflect.Int:
		return fmt.Sprintf("%d", v.i)
	case reflect.Bool:
		return fmt.Sprintf("%t", v.b)
	case reflect.Float64:
		return fmt.Sprintf("%f", v.f)
	}
	return ""
}

func (v *Value) StoreBool(b bool) {
	v.t = reflect.Bool
	v.b = b
}

func (v *Value) StoreFloat(f float64) {
	v.t = reflect.Float64
	v.f = f
}

func (v *Value) StoreInt(i int) {
	v.t = reflect.Int
	v.i = i
}

func (v *Value) StoreString(s string) {
	v.t = reflect.String
	v.s = s
}

func NewBoolValue(b bool) Value {
	value := &Value{}
	value.StoreBool(b)
	return *value
}

func NewFloatValue(f float64) Value {
	value := &Value{}
	value.StoreFloat(f)
	return *value
}

func NewIntValue(i int) Value {
	value := &Value{}
	value.StoreInt(i)
	return *value
}

func NewStringValue(s string) Value {
	value := &Value{}
	value.StoreString(s)
	return *value
}
