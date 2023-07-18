package decoder

import "strings"

const (
	// QueryBySide 展示搜索
	QueryBySide = 1
	// QueryByRef 展示引用部分
	QueryByRef = 2
	// QueryByList 展示所有
	QueryByList    = 3
	QueryByDefault = QueryByList
)

var (
	isSearchBySide = []string{
		"is_search",
	}
	isQueryByAll = []string{
		"is_list",
	}

	isQueryByRef = []string{
		"is_from_reference",
	}
)

// VerifyQueryType ?filter={"id":["646e2a2f4580dab0887c18be","646e46354580dab0887c18cb","646f9336795f29e7772afa46"]}
func VerifyQueryType(query string) int {
	if hasIncludedAll(query, isSearchBySide) {
		return QueryBySide
	} else if hasIncludedAll(query, isQueryByRef) {
		return QueryByRef
	} else if hasIncludedAll(query, isQueryByAll) {
		return QueryByList
	} else {
		return QueryByDefault
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
