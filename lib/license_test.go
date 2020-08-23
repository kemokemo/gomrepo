package gomrepo

import "testing"

func Test_getLicense(t *testing.T) {
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
		{"invalid", args{"github.com/kemokemo/hogehoge"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLicense(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLicense() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getLicense() = %v, want %v", got, tt.want)
			}
		})
	}
}
