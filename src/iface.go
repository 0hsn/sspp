package main

const (
	JSON = 0
)

type Feature struct {
	OpType                 int8
	Data, Query, DefaulVal string
}

func GetDataFromFeature(feat *Feature) string {
	switch feat.OpType {
	case JSON:
		return getJson(feat.Data, feat.Query, feat.DefaulVal)
	default:
		return ""
	}
}
