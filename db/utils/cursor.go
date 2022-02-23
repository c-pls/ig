package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Cursor struct {
	Name  string
	Value string
}

// CreateCursor return a new cursor
func CreateCursor(listCursor []Cursor) string {
	var s string
	for _, e := range listCursor {
		s += fmt.Sprintf("%s:%s?", e.Name, e.Value)
	}
	// Trim the ? suffix
	cursor := []byte(strings.TrimSuffix(s, "?"))
	return base64.StdEncoding.EncodeToString(cursor)
}

func DecodeCursor(encoded string) []Cursor {
	var res []Cursor
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(decoded), "?")
	for _, e := range list {
		re, _ := regexp.Compile("(\\w+):(.{1,})")

		group := re.FindSubmatch([]byte(e))

		res = append(res, Cursor{
			Name:  string(group[1]),
			Value: string(group[2]),
		})
	}
	return res
}
