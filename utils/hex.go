package utils

import (
	"fmt"
)

func decodeHexToString(hexString string) (string, error) {
	decodedBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", fmt.Errorf("[!] Error: %+v", err)
	}

	decodedString := fmt.Sprintf("%s", decodedBytes)

	return decodedString, nil
}
