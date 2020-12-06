package crawlers

import "strings"

func normalizeWithSymbol(s string, symbol string) string {
	if strings.Contains(s, " ") {
		return strings.Replace(s, " ", symbol, -1)
	}
	return s
}
