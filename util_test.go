package steamcommunityapi

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestFilter(t *testing.T) {
	body := "Hello from response"
	type test struct {
		name         string
		response     http.Response
		expectedErr  error
		expectedBody []byte
	}
	var tests []test

	goodStatusCode := test{
		name: "200 Status Code",
		response: http.Response{
			StatusCode: 200,
			Status:     "200 OK",
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		},
		expectedErr:  nil,
		expectedBody: []byte(body),
	}
	tests = append(tests, goodStatusCode)

	rateLimited := test{
		name: "429 Status Code",
		response: http.Response{
			StatusCode: 429,
			Status:     "429 Too Many Requests",
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		},
		expectedErr:  SteamRateLimitExceeded,
		expectedBody: []byte(body),
	}
	tests = append(tests, rateLimited)

	randomCode := test{
		name: "500 Status Code",
		response: http.Response{
			StatusCode: 500,
			Status:     "500 test",
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		},
		expectedErr:  UnexpectedSteamStatusCode,
		expectedBody: []byte(body),
	}
	tests = append(tests, randomCode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			haveBody, haveErr := filterOutBasedOnStatusCode(&tt.response)
			if !compareByteSlices(haveBody, tt.expectedBody) {
				t.Errorf("got body %v, have body %v", string(haveBody), string(tt.expectedBody))
			}
			if tt.expectedErr != haveErr {
				t.Errorf("got error '%v', have error '%v'", haveErr, tt.expectedErr)
			}
		})
	}
}

func compareByteSlices(one, two []byte) bool {
	if len(one) != len(two) {
		return false
	}
	for i, v := range one {
		if v != two[i] {
			return false
		}
	}
	return true
}
