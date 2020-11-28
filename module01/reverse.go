package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//

/*func Reverse(word string) string {
	reverse := make([]string, 0)
	for i := len(word) - 1; i >= 0; i-- {
		reverse = append(reverse, string(word[i]))
	}
	return strings.Join(reverse, "")
}*/

/*func Reverse(word string) string {
	var reverse strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		reverse.WriteByte(word[i])
	}
	return reverse.String()
}*/

func Reverse(word string) string {
	var reverse string
	for _, r := range word {
		reverse = string(r) + reverse
	}
	return reverse
}
