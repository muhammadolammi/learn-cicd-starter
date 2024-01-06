package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header    http.Header
		want      string
		wantError error
	}{
		"no header":                      {header: http.Header{}, want: "", wantError: ErrNoAuthHeaderIncluded},
		"header1":                        {header: http.Header{"Authorization": []string{"ApiKey header1"}}, want: "header1", wantError: nil},
		"malformed authorization header": {header: http.Header{"Authorization": []string{"Bearer header1"}}, want: "", wantError: ErrMalformedAuthHeader},
		"header2":                        {header: http.Header{"Authorization": []string{"ApiKey header2"}}, want: "header2", wantError: nil},
		"header3":                        {header: http.Header{"Authorization": []string{"ApiKey header3"}}, want: "header3", wantError: nil},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.header)
			if err != nil {
				if !errors.Is(tc.wantError, err) {
					t.Errorf("error is not wanted error, wanted error: %v, actual error: %v", tc.wantError, err)
				}
			}
			dif := cmp.Diff(tc.want, got)
			if dif != "" {
				t.Fatalf(dif)
			}

		})

	}
}
