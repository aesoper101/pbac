package condition

import (
	"net/http"
)

func getValuesByKey(m map[string][]string, key Key) ([]string, bool) {
	// todo: 这个地方可能会存在问题,待考证
	name := key.String()
	if values, found := m[http.CanonicalHeaderKey(name)]; found {
		return values, true
	}

	v, ok := m[name]
	return v, ok
}
