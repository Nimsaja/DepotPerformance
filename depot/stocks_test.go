package depot

import (
	"testing"
)

func TestAddStock(t *testing.T) {
	reset()

	sg := Stock{Name: "Google"}
	sa := Stock{Name: "Amazon"}
	sn := Stock{Name: "Netflix"}
	ss := Stock{Name: "Siemens"}

	if len(Get()) > 0 {
		t.Errorf("List of stocks should be empty at the beginning! Got a length of %v ", len(Get()))
	}

	Add(sg)
	if len(Get()) != 1 {
		t.Errorf("List of stocks should have one entry! Got a length of %v ", len(Get()))
	}
	Add(sa)
	if len(Get()) != 2 {
		t.Errorf("List of stocks should have two entries! Got a length of %v ", len(Get()))
	}
	Add(sn, ss)
	if len(Get()) != 4 {
		t.Errorf("List of stocks should have two entries! Got a length of %v ", len(Get()))
	}
}

func TestInitDefaultValues(t *testing.T) {
	reset()

	//should only test if there is something as these values will change over the time
	InitializeWithDefaultStocks()
	if len(Get()) == 0 {
		t.Errorf("List of stocks should have entries when initialized with default values! Got a length of %v ", len(Get()))
	}
}

func TestSum(t *testing.T) {
	reset()

	sg := Stock{Name: "Google", Count: 1, Buy: 100}
	sa := Stock{Name: "Amazon", Count: 2.5, Buy: 10.5}
	sn := Stock{Name: "Netflix", Count: 3, Buy: 0.3}
	ss := Stock{Name: "Siemens", Count: 0.2, Buy: 1000}

	Add(sg, sa, sn, ss)

	sum := SumBuy()
	if sum != 327.15 {
		t.Errorf("Expected a sum of %v. Got %v.", 327.15, sum)
	}

}
