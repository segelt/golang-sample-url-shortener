package hashutility

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_Handler(t *testing.T) {
	req_with_query, _ := http.NewRequest("GET", "/", nil)
	q := req_with_query.URL.Query()
	q.Add("endpoint", "foobar")
	req_with_query.URL.RawQuery = q.Encode()

	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "no query string given",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "No endpoint url given",
		},
		{
			name:           "happy path",
			in:             req_with_query,
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusAccepted,
			expectedBody:   fmt.Sprintf("value: %s -- added to map.", "3858f62230ac3c915f300c664312c63f"[:6]),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := NewHashPage()
			h.parseHashRequest(test.out, test.in)
			if test.out.Code != test.expectedStatus {
				t.Logf("expected: %d\ngot: %d\n", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected: %s\ngot: %s\n", test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
