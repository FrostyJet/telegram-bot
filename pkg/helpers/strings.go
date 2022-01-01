package helpers

import (
	"log"
	"regexp"
	"strings"
)

func RemoveSpecialChars(str string) string {
	re, err := regexp.Compile(`[^a-zA-Z0-9 ]+`)
	if err != nil {
		log.Fatal(err)
	}

	cleanStr := re.ReplaceAllString(str, "")

	return cleanStr
}

func PrepareQuery(query string) string {
	query = strings.Title(strings.ToLower(query))
	return strings.ReplaceAll(query, " ", "%20")
}
