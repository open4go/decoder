package decoder

import "strings"

const (
	queryBySide    = 1
	queryByRef     = 2
	queryByDefault = queryBySide
)

var (
	isSearchBySide = []string{
		"filter",
		"range",
		"sort",
	}

	isQueryByRef = []string{
		"filter",
	}
)

// http://localhost:9998/v1/auth/merchant/appcategory?filter={"id":["646e2a2f4580dab0887c18be","646e46354580dab0887c18cb","646f9336795f29e7772afa46"]}
func VerifyQueryType(query string) int {
	if hasIncludedAll(query, isSearchBySide) {
		return queryBySide
	} else if hasIncludedAll(query, isQueryByRef) {
		return queryByRef
	} else {
		return 0
	}
}

func hasIncludedAll(s string, target []string) bool {
	for _, i := range target {
		if strings.Contains(s, i) {
			continue
		}
		return false
	}
	return true
}
