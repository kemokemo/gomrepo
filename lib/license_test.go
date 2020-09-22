package gomrepo

import (
	"testing"
)

func Test_GetLicense(t *testing.T) {
	cl := NewGomClient()
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"cobra", args{"github.com/spf13/cobra"}, "Apache-2.0", false},
		{"invalid name", args{"github.com/kemokemo/hogehoge"}, "", true},
		{"invalid url", args{"kemokemo/foo"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cl.GetLicense(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLicense() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLicense() = %v, want %v", got, tt.want)
			}
		})
	}
}
