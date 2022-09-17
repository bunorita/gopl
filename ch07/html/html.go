package html

import (
	"strings"

	"github.com/bunorita/gopl/ch05/findlinks"
)

// ex7.4
func LinksFromString(s string) []string {
	return findlinks.Links(strings.NewReader(s))
}
