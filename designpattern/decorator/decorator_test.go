package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	var feature ModelFeature = &OfflineFeature{}
	input := &FeatureInput{}

	features := feature.get(input)
	assert.Equal(t, 3, len(features))
	assert.Equal(t, "offline", features["offlineFeature_key1"])
	assert.Equal(t, 0.245, features["offlineFeature_key2"])
	assert.Equal(t, 10, features["offlineFeature_key3"])

	feature = &OnlineFeature{&OfflineFeature{}}
	features = feature.get(input)
	assert.Equal(t, 6, len(features))
	assert.Equal(t, "offline", features["offlineFeature_key1"])
	assert.Equal(t, 0.245, features["offlineFeature_key2"])
	assert.Equal(t, 10, features["offlineFeature_key3"])
	assert.Equal(t, "online", features["onlineFeature_key1"])
	assert.Equal(t, 12.4545, features["onlineFeature_key2"])
	assert.Equal(t, 7, features["onlineFeature_key3"])

	feature = &CrossFeature{&OnlineFeature{&OfflineFeature{}}}
	features = feature.get(input)
	assert.Equal(t, 9, len(features))
	assert.Equal(t, "offline", features["offlineFeature_key1"])
	assert.Equal(t, 0.245, features["offlineFeature_key2"])
	assert.Equal(t, 10, features["offlineFeature_key3"])
	assert.Equal(t, "online", features["onlineFeature_key1"])
	assert.Equal(t, 12.4545, features["onlineFeature_key2"])
	assert.Equal(t, 7, features["onlineFeature_key3"])
	assert.Equal(t, "cross", features["crossFeature_key1"])
	assert.Equal(t, -0.4545, features["crossFeature_key2"])
	assert.Equal(t, 3, features["crossFeature_key3"])
}
