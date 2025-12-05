package utils

func DecodeRot13(cipherText string) (decodedString string) {
	
	var outputBuffer []byte
	
	key := byte(13 % 26 + 26) % 26

	for _, b := range []byte(cipherText) {
		newByte := b
		if 'A' <= b && b <= 'Z' {
			outputBuffer = append(outputBuffer, 'A' + (newByte - 'A' + key) % 26)
		} else if 'a' <= b && b <= 'z' {
			outputBuffer = append(outputBuffer, 'a' + (newByte - 'a' + key) % 26)
		} else {
			outputBuffer = append(outputBuffer, newByte)
		}
	}
	decodedString = string(outputBuffer)
	
	return decodedString
}
