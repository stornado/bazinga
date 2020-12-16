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
		{"", args{"//github.com", "/stornado", nil}, "//github.com/stornado", false},
		{"", args{"https://github.com/golang", "/stornado", nil}, "https://github.com/stornado", false},
		{"", args{"https://github.com", "stornado", []string{"bazinga"}}, "https://github.com/stornado/bazinga", false},
		{"", args{"https://github.com/golang", "/stornado", []string{"bazinga"}}, "https://github.com/stornado/bazinga", false},
		{"", args{"https://github.com/golang", "https://github.com/stornado", []string{"bazinga"}}, "", true},
		{"", args{"https://github.com/golang", "//github.com/stornado", []string{"bazinga"}}, "", true},
		{"", args{"https://github.com/golang", "/stornado", []string{"//github.com/stornado/bazinga"}}, "", true},
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
