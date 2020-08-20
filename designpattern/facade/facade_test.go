package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	bidder := BidderImpl{bidOptimizer: &BidOptimizer{}, predictor: &Predictor{}, detargeter: &Detargeter{}}
	bidPrice := bidder.bid(&Request{})
	assert.Equal(t, float32(50), bidPrice)
}
