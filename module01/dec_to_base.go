package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
/*func DecToBase(dec, base int) string {
	var output string
	var remainder int
	for dec > 0 {
		remainder = dec % base
		if remainder < 10 {
			output = strconv.Itoa(remainder) + output
		} else {
			switch remainder {
			case 10:
				output = "A" + output
			case 11:
				output = "B" + output
			case 12:
				output = "C" + output
			case 13:
				output = "D" + output
			case 14:
				output = "E" + output
			case 15:
				output = "F" + output
			}
		}
			dec = dec / base
	}
	return output
}
*/
func DecToBase(dec, base int) string {
	var output string
	var remainder int
	const charset = "0123456789ABCDEF"
	for dec > 0 {
		remainder = dec % base
		output = string(charset[remainder]) + output
		dec = dec / base
	}
	return output
}
