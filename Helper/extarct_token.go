package Helper

import "strings"

func ExtractToken(fullToken string) string {
	splitToken := strings.Split(fullToken, "Bearer ")
	return splitToken[1]
}
