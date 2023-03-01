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

// Factory to get data based on choice
func GetDataFromFeature(feat *Feature) string {
	switch feat.OpType {
	case JSON:
		return getJson(feat.Data, feat.Query, feat.DefaultVal)
	case XML:
		return getXml(feat.Data, feat.Query, feat.DefaultVal)
	case YAML:
		return getYaml(feat.Data, feat.Query, feat.DefaultVal)
	case TOML:
		return getToml(feat.Data, feat.Query, feat.DefaultVal)
	case INI:
		return getIni(feat.Data, feat.Query, feat.DefaultVal)
	default:
		return ""
	}
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
