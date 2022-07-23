package hashutility

import (
	"fmt"
	"gobasictinyurl/src/persistence"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestHandlers_Handler(t *testing.T) {
	req_with_query, _ := http.NewRequest("GET", "/", nil)
	q := req_with_query.URL.Query()
	q.Add("endpoint", "foobar")
	req_with_query.URL.RawQuery = q.Encode()

	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}
	defer sqldb.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 sqldb,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	// mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "url_entries" WHERE ID = $1 ORDER BY "url_entries"."id" LIMIT 1`)).
		WithArgs("foobar").
		WillReturnRows(sqlmock.NewRows([]string{"id", "value", "userid"}).
			AddRow("foobar", "7896c2", "1"))

	persistence.Instance = db

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
			expectedBody:   fmt.Sprintf("value: %s -- retrieved from map.", "7896c2"),
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
