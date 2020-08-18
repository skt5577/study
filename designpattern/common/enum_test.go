package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetModelBuildType(t *testing.T) {
	assert.Equal(t, TF_ESTIMATOR, GetModelBuildType(int(TF_ESTIMATOR)))
	assert.True(t, WORKAROUND == GetModelBuildType(2))
}
