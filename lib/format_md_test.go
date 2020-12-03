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
		{"1 row", args{[]pkginfo{{"sample", "v0.1", "MIT", nil}}}, "|ID|Version|License|\n|:---|:---|:---|\n|sample|v0.1|MIT|"},
		{"2 row", args{[]pkginfo{{"sample", "v0.1", "MIT", nil}, {"github.com/kemokemo/foo", "v0.2", "Apache-2.0", nil}}}, "|ID|Version|License|\n|:---|:---|:---|\n|github.com/kemokemo/foo|v0.2|Apache-2.0|\n|sample|v0.1|MIT|"},
		{"3 row with error", args{[]pkginfo{{"sample", "v0.1", "MIT", nil}, {"github.com/kemokemo/foo", "v0.2", "Apache-2.0", nil}, {"bar", "v0.3", "", errors.New("test")}}}, "|ID|Version|License|\n|:---|:---|:---|\n|github.com/kemokemo/foo|v0.2|Apache-2.0|\n|sample|v0.1|MIT|"},
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
