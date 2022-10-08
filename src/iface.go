package main

const (
	JSON = 0
	XML  = 1
	YAML = 2
	TOML = 3
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
	default:
		return ""
	}
}
