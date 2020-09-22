package gomrepo

import (
	"bytes"
	"testing"
)

func Test_GetLicense(t *testing.T) {
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
			got, err := GetLicense(tt.args.name)
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

var (
	moduleList = []string{
		"github.com/sirupsen/logrus v1.6.0",
		"github.com/gin-gonic/gin v1.6.3",
	}

	licenseList = `github.com/sirupsen/logrus v1.6.0 MIT
github.com/gin-gonic/gin v1.6.3 MIT
`
)

func TestPrintLicenses(t *testing.T) {
	type args struct {
		modules []string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{"normal", args{moduleList}, licenseList, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := PrintLicenses(w, tt.args.modules); (err != nil) != tt.wantErr {
				t.Errorf("printLicenses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("printLicenses() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func BenchMarkPrintLicenses(b *testing.B) {
	w := &bytes.Buffer{}
	b.ResetTimer()
	PrintLicenses(w, moduleList)
}
