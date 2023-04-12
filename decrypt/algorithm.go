package decrypt

func Nimbus(str string) string {
	descryptStr := ""

	for _, character := range str {
		asciiCode := int(character)
		descryptStr += string(asciiCode - 3)
	}

	return descryptStr
}
