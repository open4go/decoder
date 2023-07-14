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
				"filter={\"category_id\":\"646e2a2f4580dab0887c18be\", \"category_id2\":\"646e2a2f4580dab0887c18be\"}&range=[0,9]&sort=[\"id\",\"ASC\"]",
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
