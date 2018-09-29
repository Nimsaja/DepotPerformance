package depot

import "sync"

// StockValue value from a stock
type StockValue struct {
	Stock Stock
	Close float32
	Price float32
}

// StockValueList list of stock values
type StockValueList struct {
	svl          chan StockValue
	cap          int
	ready        int
	sumYesterday float32
	sumToday     float32
	//mutual exclusion - so that only on go routine can have access to StockValueList
	//block which should be mutual exclusive must be surrounded by lock and unlock
	sync.Mutex
}

// Yesterday value of stock for yesterday
func (sv StockValue) Yesterday() float32 {
	return sv.Close * sv.Stock.Count
}

// Today value of stock for today
func (sv StockValue) Today() float32 {
	return sv.Price * sv.Stock.Count
}

// New stock value list (channel with stock values)
// cap: nb of stocks
func New(cap int) *StockValueList {
	return &StockValueList{
		svl: make(chan StockValue, cap),
		cap: cap,
	}
}

// Add a stock value to channel
func (svl *StockValueList) Add(sv ...StockValue) {
	for _, el := range sv {
		svl.svl <- el
	}
}

// Done counter for get stock values are done
func (svl *StockValueList) Done() {
	svl.Lock()
	defer svl.Unlock()
	svl.ready++
	if svl.ready == svl.cap {
		close(svl.svl)
	}
}

// Wait read all stock values from channel
// and create sum today and yesterday
func (svl *StockValueList) Wait() { //[]StockValue {
	// res := make([]StockValue, svl.cap)
	// i := 0
	for sv := range svl.svl {
		// res[i] = sv
		// i++
		svl.sumYesterday += sv.Yesterday()
		svl.sumToday += sv.Today()
	}
	// return res
}

// SumYesterday sum stock values of yesterday
func (svl *StockValueList) SumYesterday() float32 {
	return svl.sumYesterday
}

// SumToday sum stock values of today
func (svl *StockValueList) SumToday() float32 {
	return svl.sumToday
}
