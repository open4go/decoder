package decoder

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
)

type UrlQuery struct {
	QueryValue url.Values
	QueryBy    int
	Translates map[string]string
}

// checkQueryFrom the origin query string
// http://localhost:9998/v1/auth/merchant/apps?filter={"category_id":"646e2a2f4580dab0887c18be"}&range=[0,9]&sort=["id","ASC"]
func (u *UrlQuery) checkQueryFrom(query string) error {
	switch VerifyQueryType(query) {
	case QueryByRef:
		u.QueryBy = QueryByRef
	case QueryBySide:
		u.QueryBy = QueryBySide
	default:
		u.QueryBy = QueryByDefault
	}
	u.Translates = make(map[string]string, 0)
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

// ?filter={"category_id":"646e3a1f4580dab0887c18c6"}&order=DESC&page=1&perPage=25&sort=created_at
func (u *UrlQuery) ExtractMapFilterAsKey2Map() []map[string]interface{} {
	filterList := make([]map[string]interface{}, 0)
	tmpFilter := make(map[string]interface{}, 0)
	filters := u.QueryValue["filter"]
	for _, f := range filters {
		err := json.Unmarshal([]byte(f), &tmpFilter)
		if err == nil {
			filterList = append(filterList, tmpFilter)
		}
	}
	return filterList
}

// ExtractFilterAsKey2Map below show an example
// ?filter={"category_id":"646e2a2f4580dab0887c18be", "category_id2":"646e2a2f4580dab0887c18be"}
// ==>
// [{"category_id":"646e2a2f4580dab0887c18be", "category_id2":"646e2a2f4580dab0887c18be"}]
func (u *UrlQuery) ExtractFilterAsKey2Map() []map[string]interface{} {
	filterList := make([]map[string]interface{}, 0)
	tmpFilter := make(map[string]interface{}, 0)
	filters := u.QueryValue["filter"]
	for _, f := range filters {
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

func (u *UrlQuery) ExtractPage() [][]string {
	sortList := make([][]string, 0)
	tmpSort := make([]string, 0)
	sortString := u.QueryValue["page"]
	for _, f := range sortString {
		err := json.Unmarshal([]byte(f), &tmpSort)
		if err == nil {
			sortList = append(sortList, tmpSort)
		}
	}
	return sortList
}

// ReferenceByMany
// http://localhost:9998?filter={"id":[["643e12a0bc01c4620486c02d"],["646e072f7867f7b362260fbb"],[]]}
// http://localhost:9998?filter={"id":["646e2a2f4580dab0887c18be","646e46354580dab0887c18cb","646f9336795f29e7772afa46"]}
// ?filter={"id":["646e06f07867f7b362260fb9"]}
func (u *UrlQuery) ReferenceByMany() []map[string][]string {
	filters := u.QueryValue["filter"]
	filterList := make([]map[string][]string, 0)
	// example: {"id":["646ef266b96c04388d10157a","646ef29eb96c04388d10157d"]}
	tmpFilter := make(map[string][][]string, 0)
	for _, f := range filters {
		err := json.Unmarshal([]byte(f), &tmpFilter)
		if err == nil {
			for k, v := range tmpFilter {

				if len(v) > 0 {
					m := map[string][]string{
						k: v[0],
					}
					filterList = append(filterList, m)
				}
			}

		}
	}
	return filterList
}

// ReferenceByOne
// http://localhost:9998?filter={"id":[["643e12a0bc01c4620486c02d"],["646e072f7867f7b362260fbb"],[]]}
// http://localhost:9998?filter={"id":["646e2a2f4580dab0887c18be","646e46354580dab0887c18cb","646f9336795f29e7772afa46"]}
// ?filter={"id":["646e06f07867f7b362260fb9"]}
func (u *UrlQuery) ReferenceByOne() []map[string][]string {
	filters := u.QueryValue["filter"]
	filterList := make([]map[string][]string, 0)
	// example: {"id":["646ef266b96c04388d10157a","646ef29eb96c04388d10157d"]}
	tmpFilter := make(map[string][]string, 0)
	for _, f := range filters {
		err := json.Unmarshal([]byte(f), &tmpFilter)
		if err == nil {
			filterList = append(filterList, tmpFilter)
		}
	}
	return filterList
}

// Translate 如果某些数据定义不能满足搜索需要
// 则需要进行翻译,例如 category_id 和 category 是需要进行转换数据库才能进行过滤
func (u *UrlQuery) Translate(a, b string) *UrlQuery {
	u.Translates[a] = b
	return u
}

// AsMongoFilter
// http://localhost:9998/v1/auth/merchant/apps?filter={"category_id":"646e2a2f4580dab0887c18be"}&range=[0,9]&sort=["id","ASC"]
func (u *UrlQuery) AsMongoFilter(fields []string, filters map[string]interface{}) bson.D {
	mongoFilters := bson.D{}
	for _, key := range fields {
		finalKey := key
		afterTranslateKey, ok := u.Translates[key]
		if ok {
			finalKey = afterTranslateKey
		}
		val, ok := filters[key]
		if ok {
			filter := bson.E{Key: finalKey, Value: val}
			mongoFilters = append(mongoFilters, filter)
		}
	}
	return mongoFilters
}

// AsMongoFilterIn
// http://localhost:9998/v1/auth/merchant/apps?filter={"category_id":"646e2a2f4580dab0887c18be"}&range=[0,9]&sort=["id","ASC"]
func (u *UrlQuery) AsMongoFilterIn(fields []map[string][]string) bson.M {
	mongoFilters := bson.M{}
	for _, i := range fields {
		for k, v := range i {
			finalKey := k
			afterTranslateKey, ok := u.Translates[k]
			if ok {
				finalKey = afterTranslateKey
			}
			// 统一转换所有引用的id
			objIds := make([]*primitive.ObjectID, 0)
			for _, i := range v {
				objID, _ := primitive.ObjectIDFromHex(i)
				objIds = append(objIds, &objID)
			}
			mongoFilters[finalKey] = bson.M{"$in": objIds}
		}
	}
	return mongoFilters
}
