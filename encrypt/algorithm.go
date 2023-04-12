package encrypt

import "fmt"

func Nimbus(str string) string {
	encryptStr := ""
	for i, c := range str {
		assiCode := int(c)
		character := string(assiCode + 3)
		encryptStr += character

		fmt.Println(i)
	}

	return encryptStr
}
