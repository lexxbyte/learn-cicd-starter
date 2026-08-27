package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "valid ApiKey header",
			headers: http.Header{"Authorization": []string{"ApiKey SOME_KEY"}},
			wantKey: "SOME_KEY",
			wantErr: false,
		},
		{
			name:    "no authorization header",
			headers: http.Header{},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "wrong scheme",
			headers: http.Header{"Authorization": []string{"Bearer SOME_KEY"}},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "missing key after scheme",
			headers: http.Header{"Authorization": []string{"ApiKey"}},
			wantKey: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if err != nil && !tt.wantErr {
				t.Fatalf("expected no error, got: %v", err)
			}
			if err == nil && tt.wantErr {
				t.Fatalf("expected an error, got nil")
			}
			if key != tt.wantKey {
				t.Fatalf("expected key %q, got %q", tt.wantKey, key)
			}
		})
	}
}
