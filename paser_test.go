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
				"filter=%7B%22id%22%3A%5B%22646ef266b96c04388d10157a%22%2C%22646ef29eb96c04388d10157d%22%5D%7D",
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
