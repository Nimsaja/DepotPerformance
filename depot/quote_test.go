package depot

import "testing"

func TestYesterday4StockValue(t *testing.T) {
	sv := StockValue{Stock: Stock{Name: "Test", Count: 0.5}, Close: 10, Price: 15}

	if sv.Yesterday() != 5 {
		t.Errorf("Expected %v, got %v", 5, sv.Yesterday())
	}
}
func TestToday4StockValue(t *testing.T) {
	sv := StockValue{Stock: Stock{Name: "Test", Count: 0.5}, Close: 10, Price: 15}

	if sv.Today() != 7.5 {
		t.Errorf("Expected %v, got %v", 7.5, sv.Yesterday())
	}
}

func TestNew(t *testing.T) {
	svl := New(3)

	if svl.cap != 3 {
		t.Errorf("Expected %v, got %v", 3, svl.cap)
	}
}

func TestAdd(t *testing.T) {
	sv := StockValue{Stock: Stock{Name: "Test", Count: 0.5}, Close: 10, Price: 15}
	svl := New(3)

	svl.Add(sv)

	if len(svl.svl) != 1 {
		t.Errorf("Expected length of %v, got %v", 1, len(svl.svl))
	}

	chanSV := <-svl.svl
	if chanSV.Stock.Name != "Test" {
		t.Errorf("Expected Stock with name %v, got %v", "Test", chanSV.Stock.Name)
	}
}

func TestAddList(t *testing.T) {
	sv1 := StockValue{Stock: Stock{Name: "Test1", Count: 0.5}, Close: 10, Price: 15}
	sv2 := StockValue{Stock: Stock{Name: "Test2", Count: 5}, Close: 107.5, Price: 150.7}
	svl := New(3)

	svl.Add(sv1, sv2)

	if len(svl.svl) != 2 {
		t.Errorf("Expected length of %v, got %v", 2, len(svl.svl))
	}

	chanSV := <-svl.svl
	if chanSV.Stock.Name != "Test1" {
		t.Errorf("Expected Stock with name %v, got %v", "Test1", chanSV.Stock.Name)
	}
}

func TestDone(t *testing.T) {
	//Don't know how :-(
}

func TestWaitAndSums(t *testing.T) {
	// sv1 := StockValue{Stock: Stock{Name: "Test1", Count: 0.5}, Close: 10, Price: 15}
	// sv2 := StockValue{Stock: Stock{Name: "Test2", Count: 5}, Close: 107.5, Price: 150.7}

	// svl := New(2)
	// svl.Add(sv1, sv2)
	// svl.Wait()

	// if svl.SumYesterday() != 542.5 {
	// 	t.Errorf("Expected a sum of %v, got %v", 542.5, svl.SumYesterday())
	// }
	// if svl.SumToday() != 761 {
	// 	t.Errorf("Expected a sum of %v, got %v", 761, svl.SumToday())
	// }
}
