package tasks

import (
	"strings"
	"testing"
)

func TestFileNameParser(t *testing.T) {
	tests := map[string]string{
		// trivial / identity
		"":     "",
		"test": "test",
		"Test": "test",
		"TEST": "test",

		// basic PascalCase
		"NewTest":    "new-test",
		"SimpleCase": "simple-case",

		// acronyms (2+ uppercase letters)
		"HTTPRequest":   "http-request",
		"HTTPServer":    "http-server",
		"JSONData":      "json-data",
		"MyXMLParser":   "my-xml-parser",
		"OAuthClient":   "o-auth-client",
		"OAuthCallback": "o-auth-callback",

		// mixed acronym + word boundaries (non-ambiguous)
		"AnotherTEST":   "another-test",
		"THISIsTest":    "this-is-test",
		"THISIsNOTTest": "this-is-not-test",

		// numbers
		"Test1":            "test1",
		"Test12":           "test12",
		"Test1Case":        "test1-case",
		"TestV2":           "test-v2",
		"TestingV2Testing": "testing-v2-testing",

		// acronym + number combos
		"HTTP2Server":    "http2-server",
		"MyHTTP2Server":  "my-http2-server",
		"OAuth2Callback": "o-auth2-callback",
		"APIResponse200": "api-response200",
		"v2APIResponse":  "v2-api-response",
	}
	for test, want := range tests {
		got := pascalCaseToKebab(test)
		if strings.Compare(got, want) != 0 {
			t.Errorf("%s: want %s, got %s\n", test, want, got)
		}
	}

}
