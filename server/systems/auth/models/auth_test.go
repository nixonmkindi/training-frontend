package models

import "testing"

func Test_checkIndex(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "valid with index", args: args{"/mkeka/data/12", "/mkeka/data"}, want: true},
		{name: "invalid with index", args: args{"/mkeka/data/12", "/mkeka/kuku"}, want: false},
		{name: "valid without index", args: args{"/mkeka/data", "/mkeka/data"}, want: true},
		{name: "invalid without index", args: args{"/mkeka/data", "/mkeka/date"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIndex(tt.args.str1, tt.args.str2); got != tt.want {
				t.Logf("str1: %v, str2: %v, output: %v\n", tt.args.str1, tt.args.str2, got)
				t.Errorf("checkIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
