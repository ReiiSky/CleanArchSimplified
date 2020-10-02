package handler

import "strings"

var knownDomainStatusCode = map[string]int{
	"internal-error": 500,
	"conflict":       409,
	"ok":             200,
	"created":        201,
	"accepted":       202,
}

// ParseDomainMessage ..
func ParseDomainMessage(message string) (int, string) {
	splited := strings.Split(message, "->")
	if len(splited) == 2 {
		x, y := knownDomainStatusCode[splited[0]]
		if y {
			return x, splited[1]
		}
	}
	return -1, strings.Join(splited, " ")
}
