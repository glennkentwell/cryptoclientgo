package cryptoclientgo

//Multiplier is the multiplier used when dealing with amounts or prices, to avoid floating point errors
const Multiplier = int64(100000000)

//ConvertToFloat converts an int64 to the corresponding float64
func ConvertToFloat(i int64) float64 {
	return float64(i) / float64(Multiplier)
}

//ConvertFromFloat converts an float64 to the corresponding int64
func ConvertFromFloat(f float64) int64 {
	return int64(f * float64(Multiplier))
}
