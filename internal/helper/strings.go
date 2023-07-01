package helper

import (
	"encoding/json"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

func IndexOfSliceString(slice []string, item string) int {
	index := -1
	for idx, itemArr := range slice {
		if itemArr == item {
			index = idx
			break
		}
	}
	return index
}

func PrettyPrint(input interface{}) string {
	jsonByte, err := json.Marshal(input)
	if err != nil {
		return err.Error()
	}
	return string(jsonByte)
}

func Normalize(str string) string {
	trans := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(trans, str)
	result = strings.ReplaceAll(result, "đ", "d")
	result = strings.ReplaceAll(result, "Đ", "D")
	return result
}
