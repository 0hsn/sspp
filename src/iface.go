package main

const (
	JSON = 0
	XML  = 1
	YAML = 2
	TOML = 3
	INI  = 4
)

// This holds parsed content
type Feature struct {
	OpType                  int8
	Data, Query, DefaultVal string
}

func interpolate(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = interpolate(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = interpolate(v)
		}
	}
	return i
}
