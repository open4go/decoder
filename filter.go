package decoder

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

// MakeFilter fields
// name=>rename
func MakeFilter(c *gin.Context, fields []string) interface{} {
	queryPath := c.Request.URL.RawQuery
	f := &UrlQuery{}
	err := f.Load(queryPath)
	if err != nil {
		return nil
	}
	filterFields := make([]string, 0)
	for _, i := range fields {
		s := strings.Split(i, "=>")
		if len(s) == 2 {
			filterFields = append(filterFields, s[0])
			f = f.Translate(s[0], s[1])
		} else {
			filterFields = append(filterFields, i)
		}
	}

	filterMap := f.ExtractFilterAsKey2Map()
	if len(filterMap) == 0 {
		return nil
	}
	filters := f.AsMongoFilter(fields, filterMap[0])
	filterInSlice := f.ExtractFilterAsKey2Map()
	if len(filterInSlice) > 0 {
		ofOneItem := filterInSlice[0]
		targetIdList, ok := ofOneItem["id"]
		if !ok {
			return nil
		}
		return ToMany(targetIdList)
	}
	return filters
}

func ToMany(v interface{}) bson.M {

	aInterface := v.([]interface{})
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	// 统一转换所有引用的id
	objIds := make([]*primitive.ObjectID, 0)
	for _, i := range aString {
		objID, _ := primitive.ObjectIDFromHex(i)
		objIds = append(objIds, &objID)
	}
	return bson.M{
		"_id": bson.M{"$in": objIds},
	}
}
