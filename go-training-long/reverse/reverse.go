package reverse

func Reverse(str string) string {
	byt := []byte(str)
	for i := 0; i < len(byt)/2; i++ {
		i2 := len(byt) - 1 - i
		byt[i], byt[i2] = byt[i2], byt[i]
	}
	return string(byt)
}
