package utils

import (
	"fmt"
	"encoding/base64"
)

func decodeBase64ToString(base64String string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return "", fmt.Errorf("[!] Error: %+v", err)
	}

	decodedString := fmt.Sprintf("%s", decodedBytes)

	return decodedString, nil
}
