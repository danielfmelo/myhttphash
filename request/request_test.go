package request_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielfmelo/myhttphash/request"
)

func TestRequestGetFuncShouldSuccess(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedBody string
		expectedErr  string
		url          string
	}{
		{
			name:         "test should get correct body",
			expectedBody: "body body",
			expectedErr:  "",
			url:          "",
		},
		{
			name:         "test should get error with broken URL",
			expectedBody: "",
			expectedErr:  `Get "https://broken-url": dial tcp:`,
			url:          "broken-url",
		},
		{
			name:         "test should get error with empty URL",
			expectedBody: "",
			expectedErr:  request.ErrParseURL.Error(),
			url:          "",
		},
		{
			name:         "test should get success without url scheme",
			expectedBody: "somebody",
			expectedErr:  "",
			url:          "google.com",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(200)
				res.Write([]byte(tc.expectedBody))
			}))
			defer func() { testServer.Close() }()
			r := request.New()
			url := testServer.URL
			if tc.expectedErr != "" || tc.url != "" {
				url = tc.url
			}
			resp, err := r.Get(url)
			if err != nil || tc.expectedErr != "" {
				if !strings.HasPrefix(err.Error(), tc.expectedErr) {
					t.Errorf("expected error prefix {%s} but got {%s}", tc.expectedErr, err.Error())
				}
			}
			if tc.expectedBody == "somebody" && string(resp) != "" {
				return
			}
			if string(resp) != string(tc.expectedBody) {
				t.Errorf("expected %s but got %s", tc.expectedBody, resp)
			}
		})
	}
}
