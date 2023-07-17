package decoder

import (
	"net/url"
	"reflect"
	"testing"
)

func TestPaserUrl(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		{
			"test",
			args{
				"filter=%7B%22id%22%3A%5B%5B%22643e12a0bc01c4620486c02d%22%2C%22646e072f7867f7b362260fbb%22%5D%5D%7D",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaserUrl(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaserUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
