package facade

type Bidder interface {
	bid(request *Request) float32
}

type BidderImpl struct {
	detargeter   *Detargeter
	predictor    *Predictor
	bidOptimizer *BidOptimizer
}

func (bidder *BidderImpl) bid(request *Request) float32 {
	bidder.detargeter.detarget(request)
	pVal := bidder.predictor.predict(request)
	return bidder.bidOptimizer.calculatePrice(request, pVal)
}

type Request struct {
}

type Detargeter struct {
}

func (detargeter *Detargeter) detarget(request *Request) {

}

type Predictor struct {
}

func (predictor *Predictor) predict(request *Request) float32 {
	return 0.5
}

type BidOptimizer struct {
}

func (optimizer *BidOptimizer) calculatePrice(request *Request, pValue float32) float32 {
	return pValue * 100
}
