package utils

import (
	"fmt"
)

func decodeRot13(cipherText string) (string, error) {
	decodedBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("[!] Error: %+v", err)
	}

	decodedString := fmt.Sprintf("%s", decodedBytes)

	return decodedString, nil
}
