package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name  string
		input string
		pass  bool
		want  string
	}{
		{name: "Good header", input: "ApiKey 8675309", pass: true, want: "8675309"},
		{name: "No key ", input: "", pass: false, want: ""},
		{name: "Invalid authorization", input: "Speak, friend", pass: false, want: ""},
	}

	for _, tc := range tests {
		head := http.Header{}
		if tc.input != "" {
			head.Add("Authorization", tc.input)
		}

		got, err := GetAPIKey(head)

		if !tc.pass && err == nil {
			t.Fatalf("%s: expected: fail, got: pass", tc.name)
		} else if tc.pass && err != nil {
			t.Fatalf("%s: expected: pass, got: fail %v", tc.name, err)
		} else if got != tc.want {
			t.Fatalf("%s: expectedP %s, got: %s", tc.name, tc.want, got)
		}
	}
}
