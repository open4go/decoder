package decoder

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/url"
)

type UrlQuery struct {
	QueryValue url.Values
	QueryBy    int
}

// checkQueryFrom the origin query string
// http://localhost:9998/v1/auth/merchant/apps?filter={"category_id":"646e2a2f4580dab0887c18be"}&range=[0,9]&sort=["id","ASC"]
func (u *UrlQuery) checkQueryFrom(query string) error {
	switch VerifyQueryType(query) {
	case queryByRef:
		u.QueryBy = queryByRef
	case queryBySide:
		u.QueryBy = queryBySide
	default:
		u.QueryBy = queryByDefault
	}
	return nil
}

// Load the origin query string
func (u *UrlQuery) Load(query string) error {
	m, err := url.ParseQuery(query)
	if err != nil {
		return err
	}
	u.QueryValue = m
	// check query type
	err = u.checkQueryFrom(query)
	if err != nil {
		return err
	}
	return nil
}

// ExtractFilterAsKey2Map below show an example
// ?filter={"category_id":"646e2a2f4580dab0887c18be", "category_id2":"646e2a2f4580dab0887c18be"}
// ==>
// [{"category_id":"646e2a2f4580dab0887c18be", "category_id2":"646e2a2f4580dab0887c18be"}]
func (u *UrlQuery) ExtractFilterAsKey2Map() []map[string]interface{} {
	filterList := make([]map[string]interface{}, 0)
	tmpFilter := make(map[string]interface{}, 0)
	filters := u.QueryValue["filter"]
	fmt.Println("filters==>", filters)
	for _, f := range filters {
		fmt.Println("f-->", f)
		err := json.Unmarshal([]byte(f), &tmpFilter)
		if err == nil {
			filterList = append(filterList, tmpFilter)
		}
	}
	return filterList
}

// ExtractRangeAsSlice
// &range=[0,9]
// ==>
// [[0,9]]
func (u *UrlQuery) ExtractRangeAsSlice() [][]int {
	rangeList := make([][]int, 0)
	tmpRange := make([]int, 0)
	rangeString := u.QueryValue["range"]
	for _, f := range rangeString {
		err := json.Unmarshal([]byte(f), &tmpRange)
		if err == nil {
			rangeList = append(rangeList, tmpRange)
		}
	}
	return rangeList
}

// ExtractSortAsSlice
// &range=[0,9]
// ==>
// [[0,9]]
func (u *UrlQuery) ExtractSortAsSlice() [][]string {
	sortList := make([][]string, 0)
	tmpSort := make([]string, 0)
	sortString := u.QueryValue["sort"]
	for _, f := range sortString {
		err := json.Unmarshal([]byte(f), &tmpSort)
		if err == nil {
			sortList = append(sortList, tmpSort)
		}
	}
	return sortList
}

// ReferenceByID
// http://localhost:9998?filter={"id":[["643e12a0bc01c4620486c02d"],["646e072f7867f7b362260fbb"],[]]}
// http://localhost:9998?filter={"id":["646e2a2f4580dab0887c18be","646e46354580dab0887c18cb","646f9336795f29e7772afa46"]}
func (u *UrlQuery) ReferenceByID() []string {
	return nil
}

// AsMongoFilter
// http://localhost:9998/v1/auth/merchant/apps?filter={"category_id":"646e2a2f4580dab0887c18be"}&range=[0,9]&sort=["id","ASC"]
func (u *UrlQuery) AsMongoFilter(fields []string, filters map[string]interface{}) bson.D {
	mongoFilters := bson.D{}
	for _, key := range fields {
		val, ok := filters[key]
		if ok {
			filter := bson.E{Key: key, Value: val}
			mongoFilters = append(mongoFilters, filter)
		}
	}
	return mongoFilters
}
