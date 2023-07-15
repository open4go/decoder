package decoder

import "go.mongodb.org/mongo-driver/bson"

func CheckFilter(q string) bson.D {
	d := UrlQuery{}
	err := d.Load(q)
	if err != nil {
		panic(err)
	}
	k2m := d.ExtractFilterAsKey2Map()
	a := []string{
		"category_id",
	}
	mgoF := d.AsMongoFilter(a, k2m[0])
	return mgoF
}

func CheckRange(q string) []int {
	d := UrlQuery{}
	err := d.Load(q)
	if err != nil {
		panic(err)
	}
	k2m := d.ExtractRangeAsSlice()
	return k2m[0]
}

func CheckSort(q string) []string {
	d := UrlQuery{}
	err := d.Load(q)
	if err != nil {
		panic(err)
	}
	k2m := d.ExtractSortAsSlice()
	return k2m[0]
}
