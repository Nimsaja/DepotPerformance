package yahoo

import (
	"time"

	"github.com/Nimsaja/DepotPerformance/depot"
)

// Spark ...
type Spark struct {
	SP HistOutput `json:"spark"`
}

// HistOutput ...
type HistOutput struct {
	Res   []HistResult `json:"result"`
	Error interface{}  `json:"error"`
}

// HistResult ...
type HistResult struct {
	Resp   []Response `json:"response"`
	Symbol string     `json:"symbol"`
}

// Response ...
type Response struct {
	// m Meta `json:"meta"`
	T []int      `json:"timestamp"`
	I Indicators `json:"indicators"`
}

// Indicators ...
type Indicators struct {
	Q []Quote `json:"quote"`
}

// Quote ...
type Quote struct {
	V []float32 `json:"close"`
}

// Data this is the data we will use in this code
type Data struct {
	S  depot.Stock
	TV map[time.Time]float32
}

// StockValue rearrange Data struct to have all values for one date in a hashmap
type StockValue struct {
	Date  time.Time
	St    depot.Stock
	Value float32
}

// CreateTimeToValueMap ...
func CreateTimeToValueMap(timestamps []int, values []float32) map[time.Time]float32 {
	var m = make(map[time.Time]float32)

	var t time.Time
	for i := 0; i < len(values); i++ {
		t = time.Unix(int64(timestamps[i]), 0)
		m[t] = values[i]
	}

	return m
}

// CreateValuesOnDateMap ...
func CreateValuesOnDateMap(data []Data) map[string][]StockValue {
	var m = make(map[string][]StockValue)

	var k string
	var exists bool
	var l []StockValue
	for _, d := range data {
		for t, v := range d.TV {
			//create key
			k = t.Format("2006-01-02")

			_, exists = m[k]
			//create new StockValueList
			if !exists {
				l = make([]StockValue, 0)
				m[k] = l
			}

			//append new StockValue
			m[k] = append(m[k], StockValue{t, d.S, v})
		}
	}

	return m
}

/** Output from yahoo as reference
{"spark":
	{"result":
		[
			{
				"symbol":"EXS2.F",
				"response":
				[
					{
						"meta":
						{
							"currency":"EUR","symbol":"EXS2.F","exchangeName":"FRA",
							"instrumentType":"ETF","firstTradeDate":1263279600,
							"gmtoffset":7200,"timezone":"CEST","exchangeTimezoneName":"Europe/Berlin",
							"chartPreviousClose":27.4,"priceHint":2,
							"currentTradingPeriod":
							{
								"pre":
								{
									"timezone":"CEST","start":1537941600,"end":1537941600,"gmtoffset":7200
								},
								"regular":
								{
									"timezone":"CEST","start":1537941600,"end":1537992000,"gmtoffset":7200
								},
								"post":
								{
									"timezone":"CEST","start":1537992000,"end":1537992000,"gmtoffset":7200
								}
							},
								"dataGranularity":"1d",
								"validRanges":["1d","5d","1mo","3mo","6mo","1y","2y","5y","10y","ytd","max"]
						},
						"timestamp":[1535349600,1535436000,1535522400,1535608800,1535695200,1535954400,1536040800,1536127200,1536213600,1536300000,1536559200,1536645600,1536732000,1536818400,1536904800,1537164000,1537250400,1537336800,1537423200,1537509600,1537952796],
						"indicators":
						{
							"quote":
							[
								{
									"close":[27.83,27.84,27.93,27.82,27.62,27.82,27.47,26.93,26.62,26.59,26.7,26.65,26.62,26.52,26.71,26.66,26.75,26.36,26.03,25.99,26.26]
								}
							],
							"adjclose":
							[
								{
									"adjclose":[27.83,27.84,27.93,27.82,27.62,27.82,27.47,26.93,26.62,26.59,26.7,26.65,26.62,26.52,26.71,26.66,26.75,26.36,26.03,25.99,26.26]
								}
							]
						}
					}
				]
			}
		],"error":null}}
*/
