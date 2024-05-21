package prices

var Prices = map[string]int64{}

func init() {
	Prices = map[string]int64{
		"CE": 10,
		"AA": 15,
		"NT": 17,
		"DE": 21,
		"YR": 23,
	}
}
