package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// tests := []struct {
	// 	name  string
	// 	input http.Header
	// 	want  string
	// 	err   error
	// }{
	// 	{
	// 		name:  "valid auth header",
	// 		input: http.Header{"Authorization": []string{"ApiKey api-key-test"}},
	// 		want:  "api-key",
	// 		err:   nil,
	// 	},
	// 	{
	// 		name:  "invalid auth header",
	// 		input: http.Header{"Auth": []string{"Apikey api-key-test"}},
	// 		want:  "",
	// 		err:   errors.New("malformed authorization header"),
	// 	},
	// 	{
	// 		name:  "no auth header",
	// 		input: http.Header{},
	// 		want:  "",
	// 		err:   ErrNoAuthHeaderIncluded,
	// 	},
	// }

	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"valid_auth_header":   {input: http.Header{"Authorization": []string{"ApiKey api-key-test"}}, want: "api-key-test"},
		"invalid_auth_header": {input: http.Header{"Auth": []string{"ApiKey api-key-test"}}, want: ""},
		"no auth header":      {input: http.Header{"User-Agent": []string{"Chrome"}}, want: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := GetAPIKey(tc.input)
			if got != tc.want {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}

}
