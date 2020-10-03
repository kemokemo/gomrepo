package gomrepo

import (
	"errors"
	"testing"
)

func Test_MD_Format(t *testing.T) {
	type args struct {
		values []pkginfo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1 row", args{[]pkginfo{{"hoge", "v0.1", "MIT", nil}}}, "|hoge|v0.1|MIT|\n"},
		{"2 row", args{[]pkginfo{{"hoge", "v0.1", "MIT", nil}, {"github.com/kemokemo/foo", "v0.2", "Apache-2.0", nil}}}, "|hoge|v0.1|MIT|\n|github.com/kemokemo/foo|v0.2|Apache-2.0|\n"},
		{"3 row with error", args{[]pkginfo{{"hoge", "v0.1", "MIT", nil}, {"github.com/kemokemo/foo", "v0.2", "Apache-2.0", nil}, {"bar", "v0.3", "", errors.New("test")}}}, "|hoge|v0.1|MIT|\n|github.com/kemokemo/foo|v0.2|Apache-2.0|\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MD.table(tt.args.values)
			if got != tt.want {
				t.Errorf("MD.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
