package core_http_response

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeader(t *testing.T) {
	tests := []struct {
		expected int
	}{
		{http.StatusOK},
		{http.StatusAccepted},
		{http.StatusInternalServerError},
		{http.StatusBadRequest},
		{http.StatusBadGateway},
		{http.StatusForbidden},
	}

	for _, tt := range tests {
		recorder := httptest.NewRecorder()
		rw := NewResponseWriter(recorder)
		rw.WriteHeader(tt.expected)
		if rw.GetStatusCodeOrPanic() != tt.expected {
			t.Fatalf("status code is not %d. got=%d",
				tt.expected, rw.GetStatusCodeOrPanic())
		}
	}
}
