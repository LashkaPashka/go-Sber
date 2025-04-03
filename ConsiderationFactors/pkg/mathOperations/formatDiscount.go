package mathoperations

func ConvertDiscount(price, numberD float64) int{
	dis := numberD * 0.01
	priDis := price*dis

	return int(price-priDis)
}