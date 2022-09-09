package main

const (
	JSON = 0
	XML  = 1
)

// This holds parsed content
type Feature struct {
	OpType                 int8
	Data, Query, DefaulVal string
}

// Factory to get data based on choice
func GetDataFromFeature(feat *Feature) string {
	switch feat.OpType {
	case JSON:
		return getJson(feat.Data, feat.Query, feat.DefaulVal)
	case XML:
		return getXml(feat.Data, feat.Query, feat.DefaulVal)
	default:
		return ""
	}
}
