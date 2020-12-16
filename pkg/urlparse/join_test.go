package urlparse

import "testing"

func TestURLJoin(t *testing.T) {
	type args struct {
		server string
		path   string
		others []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{"https://github.com", "stornado", nil}, "https://github.com/stornado", false},
		{"", args{"https://github.com", "/stornado", nil}, "https://github.com/stornado", false},
		{"", args{"https://github.com/", "/stornado", nil}, "https://github.com/stornado", false},
		{"", args{"https://github.com/stornado", "bazinga", nil}, "https://github.com/stornado/bazinga", false},
		{"", args{"github.com", "/stornado", nil}, "github.com/stornado", false},
		{"", args{"//github.com", "/stornado", nil}, "//github.com/stornado", false},
		// Fixme
		{"", args{"https://github.com/golang", "/stornado", nil}, "https://github.com/stornado", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := URLJoin(tt.args.server, tt.args.path, tt.args.others...)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLJoin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
