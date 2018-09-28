package lima

import (
	"sync"
)

// StockValue value from a stock
type StockValue struct {
	Count float32
	Close float32
	Price float32
}

// Yesterday value from stock for yesterday
func (sv StockValue) Yesterday() float32 {
	return sv.Close * sv.Count
}

// Today value from stock for today
func (sv StockValue) Today() float32 {
	return sv.Price * sv.Count
}

// StockValueList list of stock values
type StockValueList struct {
	svl          chan StockValue
	cap          int
	ready        int
	sumYesterday float32
	sumToday     float32
	sync.Mutex
}

// New stock value list (channel with stock values)
func New(cap int) *StockValueList {
	return &StockValueList{
		svl: make(chan StockValue, cap),
		cap: cap,
	}
}

// Add a stock value to channel
func (svl *StockValueList) Add(sv StockValue) {
	svl.svl <- sv
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
func (svl *StockValueList) Wait() []StockValue {
	res := make([]StockValue, svl.cap)
	i := 0
	for sv := range svl.svl {
		res[i] = sv
		i++
		svl.sumYesterday += sv.Yesterday()
		svl.sumToday += sv.Today()
	}
	return res
}

// SumYesterday sum aller stock values from yesterday
func (svl *StockValueList) SumYesterday() float32 {
	return svl.sumYesterday
}

// SumToday sum aller stock values from today
func (svl *StockValueList) SumToday() float32 {
	return svl.sumToday
}

// Gain all values from today minus yesterday (can be plus or minus)
func (svl *StockValueList) Gain() float32 {
	return svl.sumToday - svl.sumYesterday
}
