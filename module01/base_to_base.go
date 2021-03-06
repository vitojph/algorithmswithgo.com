package module01

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//   BaseToBase("E", 16, 2) => "1110"
//
func BaseToBase(input string, base, newBase int) string {
	/*	if base != 10 {
			decimalValue := BaseToDec(input, base)
			return DecToBase(decimalValue, newBase)

		}
		intDec, _ := strconv.Atoi(input)
		return DecToBase(intDec, newBase)
	*/
	return DecToBase(BaseToDec(input, base), newBase)
}
