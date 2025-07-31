package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input         http.Header
		expected      string
		expectedError error
	}

	tests := []test{
		{
			input:         http.Header{"Authorization": []string{""}},
			expected:      "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			input:         http.Header{"Authorization": []string{"authTest"}},
			expected:      "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			input:         http.Header{"Authorization": []string{" ApiKey testKey"}},
			expected:      "testKey",
			expectedError: nil,
		},
		{
			input:         http.Header{"Authorization": []string{" ApiKey testKey "}},
			expected:      "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	for _, testCase := range tests {
		result, err := GetAPIKey(testCase.input)
		if testCase.expected != result && testCase.expectedError != err {
			t.Fatalf("expected: %s, got: %v", testCase.expected, testCase.expectedError)
		}
	}
}
