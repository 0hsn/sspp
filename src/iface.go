package main

const (
	JSON = 0
)

type Feature struct {
	OpType          int8
	Data, DefaulVal string
}

func GetDataFromFeature(feat *Feature) string {
	switch feat.OpType {
	case JSON:
		return GetJsonFromFeature()
	default:
		return ""
	}
}
