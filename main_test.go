package decoder

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func TestCheckFilter(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want bson.D
	}{
		{
			"test",
			args{
				"filter=%7B%22category_id%22%3A%22646e2a2f4580dab0887c18be%22%7D&range=%5B0%2C9%5D&sort=%5B%22id%22%2C%22ASC%22%5D",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFilter(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckRange(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test",
			args{
				"filter=%7B%22category_id%22%3A%22646e2a2f4580dab0887c18be%22%7D&range=%5B0%2C9%5D&sort=%5B%22id%22%2C%22ASC%22%5D",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckRange(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
