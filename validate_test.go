package decoder

import "testing"

func TestVerifyQueryType(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 2",
			args{
				"http://localhost:9998/v1/auth/merchant/appcategory?filter=%7B%22id%22%3A%5B%22646e2a2f4580dab0887c18be%22%5D%7D",
			},
			1,
		},
		{
			"test 2",
			args{
				"\nhttp://localhost:9998/v1/auth/merchant/appcategory?filter=%7B%7D&range=%5B0%2C24%5D&sort=%5B%22id%22%2C%22DESC%22%5D",
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyQueryType(tt.args.query); got != tt.want {
				t.Errorf("VerifyQueryType() = %v, want %v", got, tt.want)
			}
		})
	}
}
