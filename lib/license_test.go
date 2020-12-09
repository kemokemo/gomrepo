package gomrepo

import (
	"bytes"
	"crypto/tls"
	"net/http"
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
		{"invalid name", args{"github.com/kemokemo/sample"}, "", true},
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

func TestGomClient_GetLicenseList(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{ServerName: "go.dev"},
	}
	client := &http.Client{
		Transport: tr,
	}

	mods, err := GetModuleList("./test-data")
	if err != nil {
		t.Errorf("failed to GetModuleList: %v", err)
		return
	}

	type fields struct {
		client *http.Client
	}
	type args struct {
		modules []string
		tf      tableFormatter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantW   string
		wantErr bool
	}{
		{"markdown normal", fields{client}, args{mods, MD}, "|ID|Version|License|\n|:---|:---|:---|\n|github.com/andybalholm/cascadia|v1.1.0|BSD-2-Clause|\n|golang.org/x/net|v0.0.0-20180218175443-cbe0f9307d01|BSD-3-Clause|\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GomClient{
				client: tt.fields.client,
			}
			w := &bytes.Buffer{}
			if err := g.GetLicenseList(w, tt.args.modules, tt.args.tf); (err != nil) != tt.wantErr {
				t.Errorf("GomClient.GetLicenseList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("GomClient.GetLicenseList() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
