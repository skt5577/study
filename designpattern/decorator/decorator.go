package decorator

type FeatureInput struct {
}
type ModelFeature interface {
	get(input *FeatureInput) map[string]interface{}
}

type OfflineFeature struct {
}

func (feature *OfflineFeature) get(input *FeatureInput) map[string]interface{} {
	features := make(map[string]interface{})
	features["offlineFeature_key1"] = "offline"
	features["offlineFeature_key2"] = 0.245
	features["offlineFeature_key3"] = 10

	return features
}

type OnlineFeature struct {
	previous ModelFeature
}

func (feature *OnlineFeature) get(input *FeatureInput) map[string]interface{} {
	features := feature.previous.get(input)
	features["onlineFeature_key1"] = "online"
	features["onlineFeature_key2"] = 12.4545
	features["onlineFeature_key3"] = 7

	return features
}

type CrossFeature struct {
	previous ModelFeature
}

func (feature *CrossFeature) get(input *FeatureInput) map[string]interface{} {
	features := feature.previous.get(input)
	features["crossFeature_key1"] = "cross"
	features["crossFeature_key2"] = -0.4545
	features["crossFeature_key3"] = 3

	return features
}
