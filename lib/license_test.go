package gomrepo

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"testing"
)

const (
	expectedMD = `|ID|Version|License|
|:---|:---|:---|
|github.com/andybalholm/cascadia|v1.1.0|BSD-2-Clause|
|golang.org/x/net|v0.0.0-20180218175443-cbe0f9307d01|BSD-3-Clause|
`

	expectedHTML = `<table>
	<thead>
		<tr>
		  <th>ID</th>
		  <th>Version</th>
		  <th>License</th>
		</tr>
	</thead>
	<tbody>
		<tr>
		  <td>github.com/andybalholm/cascadia</td>
		  <td>v1.1.0</td>
		  <td>BSD-2-Clause</td>
		</tr>
		<tr>
		  <td>golang.org/x/net</td>
		  <td>v0.0.0-20180218175443-cbe0f9307d01</td>
		  <td>BSD-3-Clause</td>
		</tr>
	</tbody>
</table>
`
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
		tf      Formatter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantW   string
		wantErr bool
	}{
		{"markdown normal", fields{client}, args{mods, MD}, expectedMD, false},
		{"HTML normal", fields{client}, args{mods, HTML}, expectedHTML, false},
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
